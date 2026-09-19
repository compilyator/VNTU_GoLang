package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type request struct {
	// Поля структури містять уже перетворені й перевірені дані.
	// Код після readRequest не повинен повторно розбирати текст.
	Count int
	Label string
}

// readRequest читає один логічний запит із довільного джерела.
// io.Reader дає змогу використовувати os.Stdin у програмі та
// strings.Reader у тестах без двох різних реалізацій.
func readRequest(reader io.Reader) (request, error) {
	scanner := bufio.NewScanner(reader)
	// Scanner читає повні рядки. Це важливо для назви з кількох слів.
	if !scanner.Scan() {
		return request{}, scanError(scanner.Err(), "очікувалося число")
	}

	// Текст спочатку очищається, потім перетворюється на число,
	// і лише після успішного перетворення перевіряється діапазон.
	countText := strings.TrimSpace(scanner.Text())
	count, err := strconv.Atoi(countText)
	if err != nil {
		return request{}, fmt.Errorf("кількість %q не є цілим числом: %w", countText, err)
	}
	if count < 1 || count > 10 {
		return request{}, fmt.Errorf("кількість має бути від 1 до 10, отримано %d", count)
	}

	if !scanner.Scan() {
		return request{}, scanError(scanner.Err(), "очікувалася назва")
	}

	// TrimSpace не змінює пробіли всередині багатослівної назви,
	// але не дозволяє прийняти рядок, що складається лише з пробілів.
	label := strings.TrimSpace(scanner.Text())
	if label == "" {
		return request{}, errors.New("назва не може бути порожньою")
	}

	return request{Count: count, Label: label}, nil
}

// scanError зберігає технічну помилку джерела через wrapping.
// Якщо джерело просто закінчилося, повертається зрозуміле пояснення
// того, яке значення програма очікувала наступним.
func scanError(err error, fallback string) error {
	if err != nil {
		return fmt.Errorf("помилка читання: %w", err)
	}
	return errors.New(fallback)
}

func main() {
	fmt.Println("Введіть кількість від 1 до 10, потім назву окремим рядком:")
	// Лише точка входу прив'язує універсальну функцію до клавіатури.
	request, err := readRequest(os.Stdin)
	if err != nil {
		// Діагностика спрямовується в stderr, а не змішується
		// зі звичайним результатом програми в stdout.
		fmt.Fprintln(os.Stderr, "Помилка:", err)
		os.Exit(1)
	}

	fmt.Printf("Прийнято: count=%d, label=%q\n", request.Count, request.Label)
}
