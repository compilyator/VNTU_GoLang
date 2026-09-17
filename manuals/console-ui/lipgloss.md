# Стилі й компонування через Lip Gloss

Офіційний репозиторій: [charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss). Для нової роботи використовуйте v2.

```powershell
go get charm.land/lipgloss/v2@latest
```

```go
import "charm.land/lipgloss/v2"

title := lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FFFFFF")).
	Background(lipgloss.Color("#5A56E0")).
	Padding(0, 1)

box := lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	Padding(1).
	Width(36)

fmt.Println(title.Render("Результати аналізу"))
fmt.Println(box.Render("Середнє значення: 18.42"))
```

Lip Gloss також має таблиці, списки, вирівнювання та гіперпосилання у сумісних терміналах. Не задавайте надмірну фіксовану ширину: перевірте вузьке вікно.

У звіті покажіть звичайний зміст, стиль, поведінку у вузькому терміналі та поясніть, чому оформлення не змішується з обчисленнями.

[До огляду](README.md)
