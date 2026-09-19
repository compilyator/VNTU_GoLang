package main

import (
	"fmt"
	"io"
	"os"
)

func renderList(writer io.Writer, title string, items []string) error {
	// Усі записи спрямовуються в переданий writer. Функція однаково
	// працює з консоллю, файлом, мережевим потоком або буфером тесту.
	if _, err := fmt.Fprintln(writer, title); err != nil {
		// Після першої помилки продовжувати запис немає сенсу:
		// викликач отримує першопричину й сам визначає реакцію.
		return err
	}
	for index, item := range items {
		// Індекс зрізу починається з нуля, а видима нумерація — з одиниці.
		if _, err := fmt.Fprintf(writer, "%d. %s\n", index+1, item); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	// Тільки main обирає реальне місце виведення. renderList не має
	// прямої залежності від os.Stdout і тому легко тестується.
	if err := renderList(os.Stdout, "Етапи перевірки", []string{"форматування", "тести", "перегляд diff"}); err != nil {
		fmt.Fprintln(os.Stderr, "Помилка виведення:", err)
		os.Exit(1)
	}
}
