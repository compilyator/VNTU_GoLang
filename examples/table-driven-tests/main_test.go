package main

import "testing"

func TestAverage(t *testing.T) {
	// Структура таблиці однакова для успішних і негативних сценаріїв.
	// Значення want задаються вручну й не залежать від average.
	tests := []struct {
		name      string
		values    []int
		want      float64
		wantError bool
	}{
		{name: "five values", values: []int{10, 20, 30, 40, 50}, want: 30},
		{name: "lower boundary", values: []int{0}, want: 0},
		{name: "upper boundary", values: []int{100}, want: 100},
		{name: "empty", values: nil, wantError: true},
		{name: "below range", values: []int{-1}, wantError: true},
		{name: "above range", values: []int{101}, wantError: true},
	}

	for _, test := range tests {
		// Ім'я рядка стає частиною назви тесту та відразу показує,
		// який саме клас вхідних даних спричинив збій.
		t.Run(test.name, func(t *testing.T) {
			got, err := average(test.values)
			if test.wantError {
				// Негативний тест успішний лише тоді, коли дані відхилено.
				if err == nil {
					t.Fatal("average() error = nil, want non-nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("average() error = %v", err)
			}
			if got != test.want {
				t.Errorf("average() = %v, want %v", got, test.want)
			}
		})
	}
}
