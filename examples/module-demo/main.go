package main

import (
	"fmt"

	"github.com/compilyator/VNTU_GoLang/examples/module-demo/greeting"
)

func main() {
	// Значення навмисно задано без введення з клавіатури: цей приклад
	// зосереджений на імпорті власного пакета, а не на роботі з консоллю.
	message, err := greeting.Build("студент")
	if err != nil {
		// Помилку з бібліотечного пакета обробляє код, який знає,
		// як повідомити про неї користувачеві.
		fmt.Println("Помилка:", err)
		return
	}

	fmt.Println(message)
}
