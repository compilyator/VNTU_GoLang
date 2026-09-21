package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"io"
	"os"
	"strings"
	"time"
)

const maxInputBytes = 1 << 20

type Item struct {
	Name     string `json:"name"`
	Quantity *int   `json:"quantity"`
	Minimum  *int   `json:"minimum"`
}

type Dataset struct {
	Title      string `json:"title"`
	Location   string `json:"location"`
	ObservedAt string `json:"observedAt"`
	Items      []Item `json:"items"`
}

// Вказівники у вхідній моделі розрізняють пропущене число та допустимий нуль.
// Шаблон отримує вже перевірені звичайні значення.
type Row struct {
	Number   int
	Name     string
	Quantity int
	Minimum  int
	Low      bool
}

type Report struct {
	Title       string
	Location    string
	ObservedAt  string
	GeneratedAt string
	Total       int
	LowCount    int
	Rows        []Row
}

func loadDataset(r io.Reader) (Dataset, error) {
	var data Dataset
	raw, err := io.ReadAll(io.LimitReader(r, maxInputBytes+1))
	if err != nil {
		return data, fmt.Errorf("читання даних: %w", err)
	}
	if len(raw) > maxInputBytes {
		return data, errors.New("файл перевищує 1 MiB")
	}
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&data); err != nil {
		return data, fmt.Errorf("декодування JSON: %w", err)
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		return data, errors.New("після набору дозволено лише пробіли, не другий документ чи текст")
	}
	return data, nil
}

func buildReport(data Dataset, now time.Time) (Report, error) {
	var report Report
	if strings.TrimSpace(data.Title) == "" || strings.TrimSpace(data.Location) == "" {
		return report, errors.New("назва й місце мають бути непорожніми")
	}
	observed, err := time.Parse(time.RFC3339, data.ObservedAt)
	if err != nil {
		return report, fmt.Errorf("час спостереження має відповідати RFC 3339: %w", err)
	}
	if len(data.Items) < 1 || len(data.Items) > 100 {
		return report, errors.New("потрібно від 1 до 100 позицій обладнання")
	}
	report = Report{
		Title: data.Title, Location: data.Location,
		ObservedAt:  observed.Format(time.RFC3339),
		GeneratedAt: now.UTC().Format(time.RFC3339),
	}
	for i, item := range data.Items {
		if strings.TrimSpace(item.Name) == "" || item.Quantity == nil || item.Minimum == nil {
			return Report{}, fmt.Errorf("позиція %d: потрібні назва, quantity і minimum", i+1)
		}
		if *item.Quantity < 0 || *item.Quantity > 10000 || *item.Minimum < 0 || *item.Minimum > 10000 {
			return Report{}, fmt.Errorf("позиція %d: кількості мають бути в межах 0–10000", i+1)
		}
		row := Row{i + 1, item.Name, *item.Quantity, *item.Minimum, *item.Quantity < *item.Minimum}
		report.Rows = append(report.Rows, row)
		report.Total += row.Quantity
		if row.Low {
			report.LowCount++
		}
	}
	return report, nil
}

func renderReport(source string, report Report) ([]byte, error) {
	tmpl, err := template.New("report").Option("missingkey=error").Parse(source)
	if err != nil {
		return nil, fmt.Errorf("розбір шаблону: %w", err)
	}
	var output bytes.Buffer
	if err := tmpl.Execute(&output, report); err != nil {
		return nil, fmt.Errorf("виконання шаблону: %w", err)
	}
	return output.Bytes(), nil
}

// Наявний файл не відкривається для запису. Невдалий новий файл прибирається.
// Це не гарантує відновлення після аварійного вимкнення живлення.
func writeNew(path string, content []byte) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o600)
	if err != nil {
		return fmt.Errorf("створення %q: %w", path, err)
	}
	n, writeErr := f.Write(content)
	if writeErr == nil && n != len(content) {
		writeErr = io.ErrShortWrite
	}
	if err := errors.Join(writeErr, f.Close()); err != nil {
		cleanupErr := os.Remove(path)
		return fmt.Errorf("запис %q (і прибирання за потреби): %w", path, errors.Join(err, cleanupErr))
	}
	return nil
}

func run(args []string, stdout, stderr io.Writer) int {
	flags := flag.NewFlagSet("html-report", flag.ContinueOnError)
	flags.SetOutput(stderr)
	input := flags.String("input", "data/sample.json", "JSON із тестовими даними")
	templatePath := flags.String("template", "report.tmpl.html", "довірений HTML-шаблон")
	output := flags.String("output", "report.html", "новий HTML-файл; перезапис заборонено")
	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return 0
		}
		return 2
	}
	if flags.NArg() != 0 || strings.TrimSpace(*input) == "" || strings.TrimSpace(*templatePath) == "" || strings.TrimSpace(*output) == "" {
		fmt.Fprintln(stderr, "потрібні непорожні шляхи; позиційні аргументи не підтримуються")
		return 2
	}
	if err := generate(*input, *templatePath, *output); err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	fmt.Fprintln(stdout, "HTML-звіт створено:", *output)
	return 0
}

func generate(input, templatePath, output string) error {
	f, err := os.Open(input)
	if err != nil {
		return fmt.Errorf("відкриття даних: %w", err)
	}
	data, readErr := loadDataset(f)
	if err := errors.Join(readErr, f.Close()); err != nil {
		return err
	}
	report, err := buildReport(data, time.Now())
	if err != nil {
		return err
	}
	// Шаблон — власний перевірений файл, а не довільне завантаження від користувача.
	source, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("читання шаблону: %w", err)
	}
	html, err := renderReport(string(source), report)
	if err != nil {
		return err
	}
	return writeNew(output, html)
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}
