# Підлекція 9А: JSON, JSONL і CSV

JSON — основне ядро; JSONL і CSV — самостійне поглиблення. Задачник отримайте від викладача. Тут тільки початкові архіви та інструкція.

- [09a-files.zip — F, входи форматів](downloads/09a-files.zip).
- [09a-module.zip — K, Go-модуль і регресійні програми](downloads/09a-module.zip).
- [SHA256SUMS.txt](downloads/SHA256SUMS.txt).

Потрібен Go 1.24 або новіший; перевірено з Go 1.25.1. Windows із PowerShell та Linux/macOS із Bash. Модуль готовий, сторонніх залежностей немає; після завантаження робота локальна.

## Перевірка та підготовка

Збережіть обидва ZIP та SHA256SUMS.txt до одного каталогу й перевірте суми перед розпакуванням.

PowerShell — порівняйте Hash із відповідними рядками SHA256SUMS.txt:
```powershell
Get-FileHash ./09a-files.zip -Algorithm SHA256
Get-FileHash ./09a-module.zip -Algorithm SHA256
cat ./SHA256SUMS.txt
```

Bash, Linux:
```bash
sha256sum -c SHA256SUMS.txt
```

Bash, macOS:
```bash
shasum -a 256 -c SHA256SUMS.txt
```

Оберіть новий, ще не наявний каталог task-001. Обидва набори мають лежати поруч.

PowerShell:
```powershell
Expand-Archive -LiteralPath ./09a-files.zip -DestinationPath ./task-001
Expand-Archive -LiteralPath ./09a-module.zip -DestinationPath ./task-001
cd ./task-001/09a-module
go run ./cmd/demo
```

Bash:
```bash
mkdir task-001
unzip 09a-files.zip -d task-001
unzip 09a-module.zip -d task-001
cd task-001/09a-module
go run ./cmd/demo
```

Початковий stdout — `READY` з LF, код 0. Входи F доступні як `../09a-files/...`. Для кожної незалежної задачі використовуйте свіжу копію. Не редагуйте входи: пробіли, лапки, відсутність кінцевого LF та неправильні документи навмисні. README у F містить розміри й SHA-256 усіх 32 файлів.

## Робота з модулем

У власному пакеті реалізуйте контракт задачі; імпорт починається з `example.com/course/formats9a/`. `model` містить спільні типи стенду, `testio` — підміни для відтворення помилок Reader/Writer без заповнення диска чи зміни системних прав.

**brokenjson, brokenjsonl і brokencsv навмисно несправні**, тільки для задач 98–100. Спочатку напишіть тест за правильним контрактом, відтворіть помилку, самостійно знайдіть причину та виправте її. Тестів спочатку немає; `[no test files]` не підтверджує правильності. Готових розв’язків, службових тестів і контрольних виправлень у ZIP немає.

PowerShell:
```powershell
go fmt ./...
go test ./... -count=1
go vet ./...
```

Bash:
```bash
go fmt ./...
go test ./... -count=1
go vet ./...
```

Для JSON перевіряйте структуру даних, а не випадковий порядок ключів. Для JSONL — також фізичні рядки й порядок записів. Для CSV — декодовані поля, header і визначені умовою точні байти. Великі цілі перевіряйте як int64 або json.Number. Не змінюйте байти через текстовий конвеєр PowerShell.

Усі дані вигадані. Формулоподібні значення перевіряйте як текст у Go-тестах; не відкривайте їх в електронних таблицях. Працюйте лише у власному навчальному каталозі або t.TempDir(). Політика spreadsheet із задачника є конкретним правилом експорту, не універсальною гарантією для всіх табличних програм.

Конспект «JSON, JSONL і CSV» та внутрішні інструкції знайдіть у своєму Google Classroom; адреси залежать від курсу.

[До навчальних архівів](../README.md)
