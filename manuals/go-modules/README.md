# Go-модулі та підключення пакетів

Інструкція пояснює, як створити тематично названий Go-проєкт, підключити зовнішній пакет, перевірити залежності та пізніше узгодити module path із GitHub-репозиторієм.

## Модуль, пакет і репозиторій

- **Модуль** — набір Go-пакетів із одним `go.mod` у корені.
- **Пакет** — Go-файли одного каталогу з однаковою інструкцією `package`.
- **Import path** — шлях конкретного пакета, наприклад `github.com/compilyator/VNTU_GoLang/labs/Lab2Data/temperature`.
- **Module path** — початок import path, записаний у директиві `module` файла `go.mod`.
- **Репозиторій** — історія файлів у Git; один репозиторій може містити один або кілька модулів, хоча для початкового проєкту краще один.

`go mod init` не створює Git-репозиторій, а `git init` не створює Go-модуль.

## Тематична назва

Створіть назву за предметом програми, а не за номером лабораторної:

```text
weather-data-analyzer
book-stats
air-quality-cli
```

PowerShell:

```powershell
mkdir weather-data-analyzer
cd weather-data-analyzer
pwd
go mod init weather-data-analyzer
```

Bash:

```bash
mkdir weather-data-analyzer
cd weather-data-analyzer
pwd
go mod init weather-data-analyzer
```

Короткий локальний module path допустимий на початку навчання. Після створення GitHub-репозиторію його потрібно замінити канонічною адресою.

## Що містить `go.mod`

Мінімальний файл:

```go.mod
module weather-data-analyzer

go 1.25
```

- `module` визначає корінь імпортів власних пакетів.
- `go` задає версію мови й поведінку інструментів, на яку орієнтується модуль.
- `require` з'являється для залежностей.
- `indirect` означає, що залежність потрібна через інший модуль, а не обов'язково є зайвою.

Перевірити активний модуль:

```powershell
go env GOMOD
go mod edit -json
```

`go env GOMOD` має показати абсолютний шлях до потрібного `go.mod`. Значення `NUL` у Windows або `/dev/null` у Unix-подібній системі означає, що команда виконується поза модулем.

## Підключення навчального генератора

Для прикладу температурного пакета:

```powershell
go get github.com/compilyator/VNTU_GoLang/labs/Lab2Data/temperature@latest
```

Складові:

- `go get` змінює залежності поточного модуля;
- повний шлях закінчується назвою конкретного пакета;
- `@latest` просить актуальну доступну версію модуля;
- якщо викладач указав версію, використовуйте її замість `latest`.

Перевірка:

```powershell
go list -m all
go list -m -json github.com/compilyator/VNTU_GoLang/labs/Lab2Data
git diff -- go.mod go.sum
```

## Імпорт і виклик

```go
package main

import (
	"fmt"

	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/temperature"
)

func main() {
	values, err := temperature.Generate(5, "Вінниця")
	if err != nil {
		fmt.Println("Помилка генерації:", err)
		return
	}
	fmt.Println(values)
}
```

Import path і package identifier пов'язані, але не тотожні: шлях знаходить каталог, а `temperature.Generate` використовує оголошене ім'я пакета.

## `go.sum`

`go.sum` містить криптографічні контрольні суми завантажених модулів і їхніх `go.mod`. Він допомагає виявити неочікувану зміну отриманої залежності.

Не потрібно:

- редагувати `go.sum` вручну;
- видаляти його через те, що він «створився автоматично»;
- додавати його до `.gitignore`.

`go.sum` слід комітити разом із відповідною зміною `go.mod`.

## `go mod tidy`

```powershell
go mod tidy
```

Команда:

- додає залежності, потрібні імпортам і тестам;
- прибирає непотрібні вимоги;
- оновлює `go.sum`;
- не замінює перегляд diff.

Після неї:

```powershell
git diff -- go.mod go.sum
go test ./...
```

## Публікація на GitHub

Після створення репозиторію `https://github.com/<username>/<repository>`:

```powershell
go mod edit -module github.com/<username>/<repository>
go mod tidy
```

Потім оновіть імпорти власних пакетів:

```go
import "github.com/<username>/<repository>/analyzer"
```

Не змінюйте імпорт навчального генератора: це інший модуль.

Перевірка:

```powershell
go list ./...
go test ./...
go vet ./...
```

## Типові помилки

### `go.mod file not found`

Перейдіть до каталогу проєкту й перевірте:

```powershell
pwd
ls
go env GOMOD
```

### `no required module provides package`

Звірте import path, виконайте `go get` для правильного пакета й потім `go mod tidy`.

### `package ... is not in std`

Часто це неправильний або надто короткий import path. Власний пакет після публікації імпортується від module path із `go.mod`.

### `import cycle not allowed`

Пакети залежать один від одного по колу. Не «лікуйте» це перейменуванням імпорту: перегляньте відповідальності й напрям залежностей.

### `imported and not used`

Go забороняє невикористані імпорти. Використайте пакет або приберіть імпорт; не додавайте фіктивне використання.

## Самоперевірка

- [ ] `go env GOMOD` показує потрібний файл.
- [ ] Каталог і module path мають тематичну назву.
- [ ] Підключено лише пакет свого варіанта.
- [ ] Помилка `Generate` перевіряється до використання значень.
- [ ] `go.mod` і `go.sum` переглянуто та збережено в Git.
- [ ] Після GitHub-публікації module path збігається з URL репозиторію.
- [ ] `go test ./...` і `go vet ./...` проходять.

## Приклад і джерела

- [Виконуваний приклад локального пакета](../../examples/module-demo/).
- [Офіційний довідник Go Modules](https://go.dev/ref/mod).
- [Tutorial: Create a Go module](https://go.dev/doc/tutorial/create-module).
- [Документація `go mod`](https://go.dev/ref/mod#go-mod-init).

[До переліку практичних інструкцій](../README.md)
