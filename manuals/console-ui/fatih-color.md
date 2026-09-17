# Кольорове виведення через `fatih/color`

Офіційний репозиторій: [github.com/fatih/color](https://github.com/fatih/color). Пакет підтримує ANSI, Windows, RGB, власні стилі та `NO_COLOR`.

```powershell
go get github.com/fatih/color@latest
```

```go
import "github.com/fatih/color"

color.Green("Опрацьовано %d значень", count)
color.Red("Помилка: %.2f поза діапазоном", value)

heading := color.New(color.FgCyan, color.Bold)
heading.Println("Результати аналізу")
```

Для доступності додавайте слова «помилка», «попередження», «успіх». Перевірте поведінку зі змінною `NO_COLOR`.

```powershell
$env:NO_COLOR='1'
go run .
```

У звіті поясніть, чому колір вимикається при неінтерактивному виведенні та як зміст зберігається без нього.

[До огляду](README.md)
