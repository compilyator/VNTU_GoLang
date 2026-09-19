package main

import (
	"bytes"
	"testing"
)

func TestRenderList(t *testing.T) {
	// bytes.Buffer накопичує всі байти, записані через io.Writer.
	// Це дозволяє перевірити результат без друку під час тесту.
	var output bytes.Buffer
	if err := renderList(&output, "Перевірка", []string{"один", "два"}); err != nil {
		t.Fatalf("renderList() error = %v", err)
	}
	const want = "Перевірка\n1. один\n2. два\n"
	// Очікуваний текст статичний і включає переведення рядків,
	// оскільки нумерація та формат є предметом цього тесту.
	if output.String() != want {
		t.Errorf("renderList() = %q, want %q", output.String(), want)
	}
}
