package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

const maxBytes = 1024

// Quantity розрізняє відсутність або null і допустимий нуль.
type Stock struct {
	Version  int    `json:"version"`
	Name     string `json:"name"`
	Quantity *int   `json:"quantity"`
}

func decode(r io.Reader) (Stock, error) {
	var stock Stock
	data, err := io.ReadAll(io.LimitReader(r, maxBytes+1))
	if err != nil {
		return stock, err
	}
	if len(data) > maxBytes {
		return stock, errors.New("перевищено 1024 байти")
	}
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&stock); err != nil {
		return stock, fmt.Errorf("JSON: %w", err)
	}
	var extra any
	if err := dec.Decode(&extra); err != io.EOF {
		if err == nil {
			return stock, errors.New("другий JSON-документ")
		}
		return stock, fmt.Errorf("дані після документа: %w", err)
	}
	if stock.Version != 1 {
		return stock, errors.New("потрібна версія 1")
	}
	if strings.TrimSpace(stock.Name) == "" {
		return stock, errors.New("потрібна назва")
	}
	if stock.Quantity == nil || *stock.Quantity < 0 || *stock.Quantity > 100 {
		return stock, errors.New("quantity має бути цілим числом від 0 до 100")
	}
	return stock, nil
}

func run() error {
	stock, err := decode(strings.NewReader(`{"version":1,"name":"Кабель","quantity":0}`))
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(stock, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
