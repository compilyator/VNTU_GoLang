// Package greeting creates small validated greeting messages.
package greeting

import (
	"errors"
	"strings"
)

// Build returns a greeting for a non-empty name.
func Build(name string) (string, error) {
	// Очищення на межі пакета не дозволяє пробілам створювати
	// візуально порожнє, але технічно непорожнє ім'я.
	cleanName := strings.TrimSpace(name)
	if cleanName == "" {
		return "", errors.New("ім'я не може бути порожнім")
	}

	// Пакет повертає готове значення, але нічого не друкує сам.
	// Місце використання вирішить, куди спрямувати результат.
	return "Вітаємо, " + cleanName + "!", nil
}
