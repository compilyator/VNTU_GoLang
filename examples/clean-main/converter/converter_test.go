package converter

import "testing"

func TestConvert(t *testing.T) {
	// Тест звертається безпосередньо до предметного пакета, тому
	// не залежить від прапорців, консолі чи форматування числа.
	result, err := Convert(2.5, "m", "cm")
	if err != nil {
		t.Fatalf("Convert() error = %v", err)
	}
	// Для цієї пари значення мають точне двійкове представлення;
	// у загальному випадку float64 порівнюють із допустимою похибкою.
	if result.OutputValue != 250 {
		t.Errorf("OutputValue = %v, want 250", result.OutputValue)
	}
}
