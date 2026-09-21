package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func readLimited(r io.Reader) ([]byte, error) {
	const limit = 8
	data, err := io.ReadAll(io.LimitReader(r, limit+1))
	if err != nil {
		return nil, err
	}
	if len(data) > limit {
		return nil, errors.New("перевищено межу 8 байтів")
	}
	return data, nil
}

// Видаляємо лише новий файл, створений цим викликом, якщо запис не завершено.
func writeNew(path string, data []byte) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o600)
	if err != nil {
		return fmt.Errorf("створення %q: %w", path, err)
	}
	n, writeErr := f.Write(data)
	if writeErr == nil && n != len(data) {
		writeErr = io.ErrShortWrite
	}
	if err := errors.Join(writeErr, f.Close()); err != nil {
		return fmt.Errorf("запис %q: %w", path, errors.Join(err, os.Remove(path)))
	}
	return nil
}

func demo() (err error) {
	dir, err := os.MkdirTemp("", "go-file-safety-")
	if err != nil {
		return err
	}
	defer func() { err = errors.Join(err, os.RemoveAll(dir)) }()
	for _, input := range []string{"12345678", "123456789"} {
		_, err := readLimited(strings.NewReader(input))
		fmt.Printf("Байтів: %d; прийнято: %t\n", len(input), err == nil)
	}
	path := filepath.Join(dir, "result.txt")
	if err := writeNew(path, []byte("перший результат")); err != nil {
		return err
	}
	err = writeNew(path, []byte("заміна"))
	if !errors.Is(err, os.ErrExist) {
		return fmt.Errorf("очікувалася відмова від перезапису: %v", err)
	}
	fmt.Println("Перезапис відхилено")
	return nil
}

func main() {
	if err := demo(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
