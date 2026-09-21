package main

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestContract(t *testing.T) {
	valid := `{"version":1,"name":"Кабель","quantity":0}`
	for name, input := range map[string]string{
		"empty": "", "null": "null", "broken": "{",
		"missing":       `{"version":1,"name":"Кабель"}`,
		"null quantity": strings.Replace(valid, `:0`, `:null`, 1),
		"type":          strings.Replace(valid, `:0`, `:"0"`, 1),
		"negative":      strings.Replace(valid, `:0`, `:-1`, 1),
		"range":         strings.Replace(valid, `:0`, `:101`, 1),
		"version":       strings.Replace(valid, `:1`, `:2`, 1),
		"name":          strings.Replace(valid, "Кабель", " ", 1),
		"unknown":       strings.Replace(valid, `"version"`, `"extra":true,"version"`, 1),
		"second":        valid + " {}", "trailing": valid + " x",
		"large": valid + strings.Repeat(" ", maxBytes),
	} {
		t.Run(name, func(t *testing.T) {
			if _, err := decode(strings.NewReader(input)); err == nil {
				t.Fatal("некоректні дані прийнято")
			}
		})
	}
	padded := valid + strings.Repeat(" ", maxBytes-len(valid))
	original, err := decode(strings.NewReader(padded))
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatal(err)
	}
	restored, err := decode(strings.NewReader(string(data)))
	if err != nil || !reflect.DeepEqual(original, restored) {
		t.Fatalf("дані змінилися: %+v, %v", restored, err)
	}
}
