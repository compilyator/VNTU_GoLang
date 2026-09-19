package greeting

import "testing"

func TestBuild(t *testing.T) {
	// Пробіли в тестових даних перевіряють не лише формування тексту,
	// а й обіцянку Build нормалізувати введене ім'я.
	message, err := Build("  Олена  ")
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}

	const want = "Вітаємо, Олена!"
	// want записано явно, а не отримано повторним викликом Build:
	// інакше тест міг би відтворити ту саму помилку, що й реалізація.
	if message != want {
		t.Errorf("Build() = %q, want %q", message, want)
	}
}
