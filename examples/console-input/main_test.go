package main

import (
	"strings"
	"testing"
)

func TestReadRequest(t *testing.T) {
	// Таблиця зберігає всі вхідні дані та очікування поруч.
	// Статичні рядки роблять кожен збій повністю відтворюваним.
	tests := []struct {
		name      string
		input     string
		wantCount int
		wantLabel string
		wantError bool
	}{
		{name: "valid multiword label", input: "5\nDemo label\n", wantCount: 5, wantLabel: "Demo label"},
		{name: "not a number", input: "five\nDemo\n", wantError: true},
		{name: "outside range", input: "11\nDemo\n", wantError: true},
		{name: "blank label", input: "5\n   \n", wantError: true},
		{name: "missing label", input: "5\n", wantError: true},
	}

	for _, test := range tests {
		// t.Run створює окремий іменований підсценарій для кожного рядка.
		t.Run(test.name, func(t *testing.T) {
			// strings.NewReader імітує введення користувача в пам'яті:
			// тест не очікує натискання клавіш і не залежить від консолі.
			got, err := readRequest(strings.NewReader(test.input))
			if test.wantError {
				// Для негативних сценаріїв сам факт відхилення даних є
				// частиною контракту; успішний результат означав би дефект.
				if err == nil {
					t.Fatal("readRequest() error = nil, want non-nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("readRequest() error = %v", err)
			}
			if got.Count != test.wantCount || got.Label != test.wantLabel {
				t.Errorf("readRequest() = %+v, want count=%d label=%q", got, test.wantCount, test.wantLabel)
			}
		})
	}
}
