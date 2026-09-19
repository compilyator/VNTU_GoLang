package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	// Кожен рядок описує весь зовнішній контракт одного запуску:
	// аргументи, код завершення та суттєві фрагменти двох потоків.
	tests := []struct {
		name       string
		args       []string
		wantCode   int
		wantStdout string
		wantStderr string
	}{
		{name: "convert", args: []string{"length", "--value", "250", "--from", "cm", "--to", "m"}, wantCode: 0, wantStdout: "2.50 m"},
		{name: "missing units", args: []string{"length", "--value", "2"}, wantCode: 2, wantStderr: "обов'язковими"},
		{name: "unknown command", args: []string{"unknown"}, wantCode: 2, wantStderr: "невідома команда"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Окремі буфери дають змогу переконатися, що звичайний
			// результат і діагностика не змішуються.
			var stdout bytes.Buffer
			var stderr bytes.Buffer
			code := run(test.args, &stdout, &stderr)
			if code != test.wantCode {
				t.Errorf("run() code = %d, want %d", code, test.wantCode)
			}
			// Для очікуваного повідомлення перевіряється суттєвий фрагмент,
			// щоб тест не став крихким через несуттєве оформлення. Якщо
			// повідомлення не очікується, потік має залишитися порожнім.
			if test.wantStdout != "" && !strings.Contains(stdout.String(), test.wantStdout) {
				t.Errorf("stdout = %q, want substring %q", stdout.String(), test.wantStdout)
			}
			if test.wantStdout == "" && stdout.Len() != 0 {
				t.Errorf("stdout = %q, want empty", stdout.String())
			}
			if test.wantStderr != "" && !strings.Contains(stderr.String(), test.wantStderr) {
				t.Errorf("stderr = %q, want substring %q", stderr.String(), test.wantStderr)
			}
			if test.wantStderr == "" && stderr.Len() != 0 {
				t.Errorf("stderr = %q, want empty", stderr.String())
			}
		})
	}
}
