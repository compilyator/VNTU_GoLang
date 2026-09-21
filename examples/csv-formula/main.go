package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func spreadsheetText(value string) string {
	if value == "" {
		return value
	}
	if strings.ContainsRune("=+-@\t\r", rune(value[0])) {
		return "'" + value
	}
	return value
}

func main() {
	writer := csv.NewWriter(os.Stdout)
	for _, value := range []string{"звичайний текст", "=1+1", "@SUM(A1:A2)"} {
		if err := writer.Write([]string{spreadsheetText(value)}); err != nil {
			fmt.Fprintln(os.Stderr, "Помилка запису:", err)
			return
		}
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		fmt.Fprintln(os.Stderr, "Помилка flush:", err)
	}
}
