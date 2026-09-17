# Повноцінний TUI через Bubble Tea

Офіційний підручник: [charm.land/bubbletea](https://charm.land/bubbletea). Bubble Tea v2 реалізує подієву модель `Model–Update–View` і призначений студентам, які вже розуміють структури, методи й інтерфейси.

```powershell
go get charm.land/bubbletea/v2@latest
```

Мінімальний каркас:

```go
import tea "charm.land/bubbletea/v2"

type model struct {
	values []float64
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) View() tea.View {
	return tea.NewView("Консольний аналізатор")
}

func runUI() error {
	_, err := tea.NewProgram(model{}).Run()
	return err
}
```

Для бонусу потрібні реальні взаємодії: введення кількості, запуск генерації, перегляд значень і статистики, вихід через клавішу та обробка помилок. Статичний екран не є повноцінним TUI.

Bubble Tea керує стандартним виведенням, тому налагоджувальні журнали краще спрямовувати у файл. У лабораторній № 2 весь код усе ще має залишатися в `main.go`.

[До огляду](README.md)
