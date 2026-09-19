package main

import "fmt"

func average(values []int) (float64, error) {
	// Порожній набір не має визначеного середнього в межах цього
	// контракту, тому повертається помилка, а не умовний нуль.
	if len(values) == 0 {
		return 0, fmt.Errorf("набір не може бути порожнім")
	}

	sum := 0
	for _, value := range values {
		// Валідація виконується до додавання: результат ніколи не
		// обчислюється на частково некоректному наборі.
		if value < 0 || value > 100 {
			return 0, fmt.Errorf("значення %d поза діапазоном 0..100", value)
		}
		sum += value
	}

	// Явне перетворення обох операндів запобігає цілочисельному діленню.
	return float64(sum) / float64(len(values)), nil
}

func main() {
	// Малий статичний набір дозволяє вручну перевірити очікуване
	// значення: (10 + 20 + 30 + 40 + 50) / 5 = 30.
	result, err := average([]int{10, 20, 30, 40, 50})
	if err != nil {
		fmt.Println("Помилка:", err)
		return
	}
	fmt.Printf("Середнє: %.2f\n", result)
}
