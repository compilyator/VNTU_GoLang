package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

type Event struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

func decodeLine(line []byte) (Event, error) {
	var event Event
	decoder := json.NewDecoder(bytes.NewReader(line))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&event); err != nil {
		return event, err
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		if err == nil {
			return event, errors.New("рядок містить другий JSON value")
		}
		return event, fmt.Errorf("дані після JSON value: %w", err)
	}
	return event, nil
}

func main() {
	input := "{\"id\":1,\"status\":\"ok\"}\n" +
		"{\"id\":2,\"status\":\"rejected\"}\n" +
		"{\"id\":3,\"status\":\"ok\",\"extra\":true}\n"

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Buffer(make([]byte, 1024), 64*1024)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		event, err := decodeLine(scanner.Bytes())
		if err != nil {
			fmt.Printf("рядок %d: %v\n", lineNumber, err)
			return
		}
		fmt.Printf("рядок %d: id=%d status=%s\n", lineNumber, event.ID, event.Status)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Помилка потоку:", err)
	}
}
