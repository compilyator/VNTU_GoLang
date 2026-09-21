package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func fixture(t *testing.T) string {
	t.Helper()
	raw, err := os.ReadFile("data/sample.json")
	if err != nil {
		t.Fatal(err)
	}
	return string(raw)
}

func TestAnalysisAndEscaping(t *testing.T) {
	data, err := loadDataset(strings.NewReader(fixture(t)))
	if err != nil {
		t.Fatal(err)
	}
	data.Title = `<script>document.title='НЕБЕЗПЕЧНО'</script>`
	now := time.Date(2026, 9, 21, 12, 0, 0, 0, time.UTC)
	report, err := buildReport(data, now)
	if err != nil {
		t.Fatal(err)
	}
	if report.Total != 15 || report.LowCount != 2 || len(report.Rows) != 3 {
		t.Fatalf("неправильний аналіз: %+v", report)
	}
	if report.GeneratedAt != "2026-09-21T12:00:00Z" {
		t.Fatal(report.GeneratedAt)
	}
	source, err := os.ReadFile("report.tmpl.html")
	if err != nil {
		t.Fatal(err)
	}
	html, err := renderReport(string(source), report)
	if err != nil {
		t.Fatal(err)
	}
	for _, expected := range []string{"&lt;script&gt;", "Набір &lt;навчальний&gt; &amp; запасний", "Поповнити", ">15<", ">2<"} {
		if !bytes.Contains(html, []byte(expected)) {
			t.Errorf("відсутній текст %q", expected)
		}
	}
	if bytes.Contains(html, []byte(data.Title)) {
		t.Fatal("недовірена розмітка не екранована")
	}
}

func TestInvalidData(t *testing.T) {
	valid := fixture(t)
	cases := map[string]string{
		"empty": "", "broken": "{", "null": "null",
		"second": valid + " {}", "trailing": valid + " сторонній текст",
		"unknown":          strings.Replace(valid, `"title":`, `"extra": 1, "title":`, 1),
		"missing quantity": strings.Replace(valid, `"quantity": 12,`, "", 1),
		"negative":         strings.Replace(valid, `"quantity": 12`, `"quantity": -1`, 1),
		"wrong type":       strings.Replace(valid, `"quantity": 12`, `"quantity": "12"`, 1),
		"bad date":         strings.Replace(valid, "2026-09-20T10:00:00+03:00", "учора", 1),
		"too large":        strings.Repeat(" ", maxInputBytes+1),
	}
	for name, raw := range cases {
		t.Run(name, func(t *testing.T) {
			data, err := loadDataset(strings.NewReader(raw))
			if err == nil {
				_, err = buildReport(data, time.Time{})
			}
			if err == nil {
				t.Fatal("некоректні дані прийнято")
			}
		})
	}
}

func TestOutputAndTemplateFailures(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "report.html")
	if err := writeNew(path, []byte("попередній вміст")); err != nil {
		t.Fatal(err)
	}
	if err := generate("data/sample.json", "report.tmpl.html", path); err == nil {
		t.Fatal("дозволено перезапис")
	}
	raw, err := os.ReadFile(path)
	if err != nil || string(raw) != "попередній вміст" {
		t.Fatalf("файл пошкоджено: %q, %v", raw, err)
	}
	for _, source := range []string{"{{", "{{.UnknownField}}"} {
		templatePath := filepath.Join(dir, "bad.tmpl.html")
		if err := os.WriteFile(templatePath, []byte(source), 0o600); err != nil {
			t.Fatal(err)
		}
		output := filepath.Join(dir, "absent.html")
		if err := generate("data/sample.json", templatePath, output); err == nil {
			t.Fatal("помилку шаблону пропущено")
		}
		if _, err := os.Stat(output); !os.IsNotExist(err) {
			t.Fatal("створено неповний вихід", err)
		}
	}
	if err := generate("data/sample.json", "report.tmpl.html", filepath.Join(dir, "missing", "out.html")); err == nil {
		t.Fatal("пропущено помилку шляху")
	}
}

func TestCLI(t *testing.T) {
	var out, diagnostics bytes.Buffer
	if code := run([]string{"--input", ""}, &out, &diagnostics); code != 2 {
		t.Fatal(code)
	}
	path := filepath.Join(t.TempDir(), "result.html")
	if code := run([]string{"--output", path}, &out, &diagnostics); code != 0 {
		t.Fatal(code, diagnostics.String())
	}
	if code := run([]string{"--output", path}, &out, &diagnostics); code != 1 {
		t.Fatal(code)
	}
}
