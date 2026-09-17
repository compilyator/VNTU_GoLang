# Командний інтерфейс через Cobra

Офіційний репозиторій: [spf13/cobra](https://github.com/spf13/cobra). Cobra створює CLI з командами, підкомандами, аргументами, прапорцями, `--help` та автодоповненням.

```powershell
go get github.com/spf13/cobra@latest
```

Ідея інтерфейсу:

```text
analyzer generate --count 20 --seed 42
analyzer analyze --show-duplicates
analyzer version
```

Скорочена коренева команда:

```go
var rootCmd = &cobra.Command{
	Use:   "analyzer",
	Short: "Аналізатор числових спостережень",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAnalysis()
	},
}
```

Cobra природно потребує функцій і часто кількох файлів, тому основне місце для нього — лабораторна № 3. У № 2 він не повинен замінювати обов’язкове консольне введення і пояснення базових перевірок.

[До огляду](README.md)
