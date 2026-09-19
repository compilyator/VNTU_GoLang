// Package console renders converter results for a text console.
package console

import (
	"fmt"
	"io"

	"github.com/compilyator/VNTU_GoLang/examples/clean-main/converter"
)

// Render writes a human-readable conversion result.
func Render(writer io.Writer, result converter.Result) error {
	// io.Writer не прив'язує форматування до os.Stdout: викликач може
	// передати консоль, файл або bytes.Buffer у тесті.
	_, err := fmt.Fprintf(writer, "%.2f %s = %.2f %s\n", result.InputValue, result.InputUnit, result.OutputValue, result.OutputUnit)
	// Помилка запису повертається викликачеві. Пакет представлення
	// не вирішує самостійно, чи завершувати через неї процес.
	return err
}
