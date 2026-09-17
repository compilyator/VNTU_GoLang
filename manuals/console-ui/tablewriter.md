# Таблиці через `tablewriter`

Офіційний репозиторій: [github.com/olekukonko/tablewriter](https://github.com/olekukonko/tablewriter). У мануалі використовується актуальна гілка v1; старі приклади для `v0.0.5` можуть мати інший API.

```powershell
go get github.com/olekukonko/tablewriter@latest
```

```go
import (
	"os"

	"github.com/olekukonko/tablewriter"
)

table := tablewriter.NewWriter(os.Stdout)
table.Header("Категорія", "Кількість")
table.Append([]string{"Низькі", "4"})
table.Append([]string{"Середні", "9"})
table.Append([]string{"Високі", "2"})
table.Render()
```

Значення потрібно обчислити власною логікою до передавання таблиці. Бібліотека лише форматує результат.

У звіті зафіксуйте версію через `go list -m github.com/olekukonko/tablewriter` і покажіть, що порядок рядків відповідає умові, а не порядку обходу `map`.

[До огляду](README.md)
