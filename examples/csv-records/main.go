package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	input := "name,comment\n" +
		"Марія,\"значення, перевірено\"\n" +
		"Іван,\"перший рядок\nдругий рядок\"\n"

	reader := csv.NewReader(strings.NewReader(input))
	reader.FieldsPerRecord = 2
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Помилка CSV:", err)
			return
		}
		fmt.Printf("%q\n", record)
	}

	var output bytes.Buffer
	writer := csv.NewWriter(&output)
	if err := writer.Write([]string{"name", "comment"}); err != nil {
		fmt.Println("Помилка запису:", err)
		return
	}
	if err := writer.Write([]string{"Олена", "готово, перевірено"}); err != nil {
		fmt.Println("Помилка запису:", err)
		return
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		fmt.Println("Помилка flush:", err)
		return
	}
	fmt.Print(output.String())
}
