package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

func run(args []string, stdout, stderr io.Writer) int {
	// run не читає os.Args і не завершує процес самостійно. Явні
	// аргументи та потоки роблять зовнішню поведінку зручною для тестів.
	if len(args) == 0 {
		fmt.Fprintln(stderr, "Помилка: очікується команда length або version")
		return 2
	}

	switch args[0] {
	case "length":
		// Перший аргумент уже використано як назву підкоманди.
		// Конкретна команда отримує тільки власні прапорці.
		return runLength(args[1:], stdout, stderr)
	case "version":
		fmt.Fprintln(stdout, "converter 1.0.0")
		return 0
	case "help", "-h", "--help":
		printUsage(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "Помилка: невідома команда %q\n", args[0])
		return 2
	}
}

func runLength(args []string, stdout, stderr io.Writer) int {
	// Окремий FlagSet ізолює прапорці підкоманди length від інших
	// підкоманд. ContinueOnError повертає помилку замість os.Exit.
	flags := flag.NewFlagSet("length", flag.ContinueOnError)
	// Стандартні повідомлення пакета flag теж мають потрапити в stderr,
	// переданий викликом, а не в глобальний потік процесу.
	flags.SetOutput(stderr)
	value := flags.Float64("value", 0, "числове значення")
	from := flags.String("from", "", "початкова одиниця: m або cm")
	to := flags.String("to", "", "цільова одиниця: m або cm")
	flags.Usage = func() {
		fmt.Fprintln(stderr, "Використання: converter length --value NUMBER --from UNIT --to UNIT")
		flags.PrintDefaults()
	}

	if err := flags.Parse(args); err != nil {
		return 2
	}
	// Ця команда не підтримує позиційних аргументів після прапорців.
	// Мовчазне ігнорування могло б приховати друкарську помилку.
	if flags.NArg() != 0 {
		fmt.Fprintf(stderr, "Помилка: зайві аргументи: %v\n", flags.Args())
		return 2
	}

	// Розбір CLI завершено. Подальша функція працює зі звичайними
	// Go-значеннями й нічого не знає про прапорці або потоки.
	result, err := convertLength(*value, *from, *to)
	if err != nil {
		fmt.Fprintln(stderr, "Помилка:", err)
		return 2
	}

	fmt.Fprintf(stdout, "%.2f %s = %.2f %s\n", *value, *from, result, *to)
	return 0
}

func convertLength(value float64, from, to string) (float64, error) {
	// Порожні одиниці означають, що обов'язкові текстові прапорці
	// не були задані або отримали порожнє значення.
	if from == "" || to == "" {
		return 0, errors.New("прапорці --from і --to є обов'язковими")
	}
	if value < 0 {
		return 0, errors.New("довжина не може бути від'ємною")
	}
	if from == to && (from == "m" || from == "cm") {
		return value, nil
	}
	if from == "m" && to == "cm" {
		return value * 100, nil
	}
	if from == "cm" && to == "m" {
		return value / 100, nil
	}
	return 0, fmt.Errorf("непідтримуване перетворення %q -> %q", from, to)
}

func printUsage(writer io.Writer) {
	// Довідка приймає writer, тому однаково працює з stdout і буфером.
	fmt.Fprintln(writer, "Команди:")
	fmt.Fprintln(writer, "  length   перетворити метри й сантиметри")
	fmt.Fprintln(writer, "  version  показати версію")
}

func main() {
	// Лише межа з операційною системою використовує os.Args, стандартні
	// потоки та os.Exit. Решту поведінки можна викликати як функцію.
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}
