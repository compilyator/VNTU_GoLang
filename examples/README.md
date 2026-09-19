# Навчальні приклади Go

Приклади демонструють окремі техніки для лабораторних робіт № 2–3, але не є готовим каркасом жодного індивідуального варіанта. У них використано нейтральні предметні області: привітання, перетворення довжини, калібрувальні вимірювання й текстові звіти.

Усі каталоги входять до одного Go-модуля. Із каталогу `examples` виконайте:

```powershell
gofmt -w .
go test ./...
go vet ./...
```

## Приклади

- [`module-demo`](module-demo/) — локальний пакет і імпорт за module path.
- [`console-input`](console-input/) — читання двох рядків, parsing і валідація без залежності від `os.Stdin` у тестах.
- [`flag-subcommands`](flag-subcommands/) — підкоманди, `flag.FlagSet`, потоки та коди завершення.
- [`table-driven-tests`](table-driven-tests/) — табличні позитивні, негативні й граничні тести.
- [`clean-main`](clean-main/) — розділення `main`, предметної логіки та представлення.
- [`writer-output`](writer-output/) — формування результату через `io.Writer` і перевірка через `bytes.Buffer`.

Кожен приклад запускається з кореня цього модуля:

```powershell
go run ./module-demo
go run ./console-input
go run ./flag-subcommands length -value 250 -from cm -to m
go run ./table-driven-tests
go run ./clean-main -value 2.5 -from m -to cm
go run ./writer-output
```

Перед перенесенням ідеї до власної роботи поясніть, яку відповідальність має кожна функція та чому її контракт підходить вашій предметній області.

[До головної сторінки публічних матеріалів](../README.md)
