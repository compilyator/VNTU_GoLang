package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestReadBoundary(t *testing.T) {
	for _, n := range []int{0, 8, 9} {
		_, err := readLimited(strings.NewReader(strings.Repeat("x", n)))
		if (err != nil) != (n > 8) {
			t.Fatalf("байтів %d: %v", n, err)
		}
	}
}

func TestPreserveExisting(t *testing.T) {
	path := filepath.Join(t.TempDir(), "дані з пробілами.txt")
	if err := writeNew(path, []byte("оригінал")); err != nil {
		t.Fatal(err)
	}
	if err := writeNew(path, []byte("заміна")); !errors.Is(err, os.ErrExist) {
		t.Fatal(err)
	}
	data, err := os.ReadFile(path)
	if err != nil || string(data) != "оригінал" {
		t.Fatalf("%q: %v", data, err)
	}
	if err := writeNew(filepath.Join(t.TempDir(), "missing", "out"), nil); err == nil {
		t.Fatal("відсутній каталог прийнято")
	}
}
