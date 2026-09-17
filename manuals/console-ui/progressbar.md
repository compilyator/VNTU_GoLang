# Індикатор прогресу через `progressbar`

Офіційний репозиторій: [schollz/progressbar](https://github.com/schollz/progressbar). Пакет v3 надає однорядковий потокобезпечний progress bar.

```powershell
go get github.com/schollz/progressbar/v3@latest
```

```go
bar := progressbar.Default(int64(len(values)))
for _, value := range values {
	process(value)
	if err := bar.Add(1); err != nil {
		fmt.Println("Помилка індикатора:", err)
		break
	}
}
```

Для 5–100 чисел аналіз завершується майже миттєво, тому progress bar у лабораторній № 2 зазвичай не потрібний. Заборонено додавати `time.Sleep` лише для демонстрації анімації.

Доречні сценарії в роботах № 4–5: читання великого файла, серія HTTP-запитів, завантаження відповіді або тривале пакетне опрацювання.

[До огляду](README.md)
