# `promptui`: старіший інтерактивний інтерфейс

Репозиторій: [manifoldco/promptui](https://github.com/manifoldco/promptui). Пакет створює текстові prompt-и, перевірку та вибір зі списку.

```powershell
go get github.com/manifoldco/promptui@latest
```

```go
prompt := promptui.Prompt{
	Label: "Кількість значень",
	Validate: func(input string) error {
		// Перетворення та перевірка 5–100.
		return nil
	},
}

value, err := prompt.Run()
```

Цей пакет використовує старішу екосистему `readline`. Для нової лабораторної рекомендовано [Huh](huh.md), який має сучасні форми, теми й доступний режим. `promptui` залишено для порівняння або підтримки наявного коду; окремого бонусу лише за його вибір немає.

[До огляду](README.md)
