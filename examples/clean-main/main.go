package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/compilyator/VNTU_GoLang/examples/clean-main/console"
	"github.com/compilyator/VNTU_GoLang/examples/clean-main/converter"
)

func main() {
	// Точка входу відповідає лише за зовнішній CLI-контракт:
	// отримує прапорці, викликає потрібні пакети й обробляє помилки.
	value := flag.Float64("value", 0, "числове значення")
	from := flag.String("from", "", "початкова одиниця")
	to := flag.String("to", "", "цільова одиниця")
	flag.Parse()

	// Формула та правила допустимих одиниць приховані в предметному
	// пакеті. main не дублює їх умовами чи обчисленнями.
	conversion, err := converter.Convert(*value, *from, *to)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Помилка:", err)
		os.Exit(1)
	}

	// Пакет представлення отримує вже готовий предметний результат.
	// Передавання os.Stdout тут лишає console незалежним від глобального потоку.
	if err := console.Render(os.Stdout, conversion); err != nil {
		fmt.Fprintln(os.Stderr, "Помилка виведення:", err)
		os.Exit(1)
	}
}
