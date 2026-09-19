package console

import (
	"bytes"
	"testing"

	"github.com/compilyator/VNTU_GoLang/examples/clean-main/converter"
)

func TestRender(t *testing.T) {
	// Буфер реалізує io.Writer і зберігає результат у пам'яті,
	// тому тест не перехоплює глобальний stdout.
	var output bytes.Buffer
	// Предметний результат задається статично: цей тест перевіряє
	// тільки подання, а не повторно алгоритм перетворення.
	result := converter.Result{InputValue: 2.5, InputUnit: "m", OutputValue: 250, OutputUnit: "cm"}
	if err := Render(&output, result); err != nil {
		t.Fatalf("Render() error = %v", err)
	}
	const want = "2.50 m = 250.00 cm\n"
	// Точне порівняння виправдане, бо формат тексту є контрактом Render.
	if output.String() != want {
		t.Errorf("Render() = %q, want %q", output.String(), want)
	}
}
