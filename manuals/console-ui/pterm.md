# Комплексний консольний звіт через PTerm

Офіційні матеріали: [pterm/pterm](https://github.com/pterm/pterm) і [docs.pterm.sh](https://docs.pterm.sh/). PTerm надає повідомлення, таблиці, дерева, діаграми, progress bar, spinner та інтерактивні поля.

```powershell
go get github.com/pterm/pterm@latest
```

```go
import "github.com/pterm/pterm"

pterm.Info.Println("Аналіз набору")
pterm.Success.Printf("Опрацьовано %d значень\n", count)

data := pterm.TableData{
	{"Показник", "Значення"},
	{"Мінімум", "4.20"},
	{"Максимум", "31.70"},
	{"Середнє", "17.84"},
}

if err := pterm.DefaultTable.WithHasHeader().WithData(data).Render(); err != nil {
	fmt.Println("Не вдалося вивести таблицю:", err)
}
```

Не запускайте spinner або progress bar для миттєвої операції й не додавайте штучної затримки. Для № 2 найкраще використати повідомлення, таблицю або просту діаграму категорій.

[До огляду](README.md)
