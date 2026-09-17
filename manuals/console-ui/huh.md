# Інтерактивні форми через Huh

Офіційний репозиторій: [charmbracelet/huh](https://github.com/charmbracelet/huh). Huh підтримує `Input`, `Text`, `Select`, `MultiSelect`, `Confirm`, перевірку, теми й доступний режим. Для нової роботи використовуйте v2.

```powershell
go get charm.land/huh/v2@latest
```

```go
import (
	"errors"
	"fmt"
	"strconv"

	"charm.land/huh/v2"
)

var countText string
err := huh.NewInput().
	Title("Кількість спостережень").
	Value(&countText).
	Validate(func(value string) error {
		count, err := strconv.Atoi(value)
		if err != nil {
			return errors.New("введіть ціле число")
		}
		if count < 5 || count > 100 {
			return errors.New("допустима кількість: від 5 до 100")
		}
		return nil
	}).
	Run()
if err != nil {
	fmt.Println("Не вдалося отримати значення:", err)
	return
}

count, err := strconv.Atoi(countText)
if err != nil {
	fmt.Println("Не вдалося перетворити кількість:", err)
	return
}
fmt.Println("Кількість:", count)
```

API конкретної версії потрібно звірити з документацією після `go get`. Навіть якщо бібліотека викликає валідатор, студент має пояснити перевірку меж.

Для доступності передбачте режим без повноекранного TUI. Для форми він вмикається через `form.WithAccessible(true)`; значення краще брати з налаштування або змінної середовища, щоб режим обирав користувач.

У звіті покажіть правильне, неправильне та граничне введення.

[До огляду](README.md)
