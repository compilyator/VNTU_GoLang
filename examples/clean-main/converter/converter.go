// Package converter contains unit-conversion domain logic.
package converter

import "fmt"

// Result describes one completed conversion.
type Result struct {
	// Result зберігає як вихідні, так і обчислені дані, потрібні
	// зовнішньому шару для подання без повторного обчислення.
	InputValue  float64
	InputUnit   string
	OutputValue float64
	OutputUnit  string
}

// Convert converts between metres and centimetres.
func Convert(value float64, from, to string) (Result, error) {
	// Валідація є частиною предметного контракту, тому розташована
	// тут, а не лише в CLI. Інший клієнт пакета отримає ті самі правила.
	if value < 0 {
		return Result{}, fmt.Errorf("значення %.2f не може бути від'ємним", value)
	}

	// Спільні поля заповнюються один раз. Кожна підтримувана гілка
	// нижче визначає лише обчислене OutputValue.
	result := Result{InputValue: value, InputUnit: from, OutputUnit: to}
	switch {
	case from == "m" && to == "cm":
		result.OutputValue = value * 100
	case from == "cm" && to == "m":
		result.OutputValue = value / 100
	case from == to && (from == "m" || from == "cm"):
		result.OutputValue = value
	default:
		// Невідому пару не варто мовчки трактувати як тотожне
		// перетворення: явна помилка не приховує неправильний виклик.
		return Result{}, fmt.Errorf("непідтримуване перетворення %q -> %q", from, to)
	}

	return result, nil
}
