package releaseassets_test

import (
	"reflect"
	"testing"

	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/releaseassets"
)

func TestContract(t *testing.T) {
	for _, count := range []int{5, 20, 100} {
		first, err := releaseassets.GenerateWithSeed(count, "cli/cli", 42)
		if err != nil {
			t.Fatal(err)
		}
		second, err := releaseassets.GenerateWithSeed(count, " CLI/CLI ", 42)
		if err != nil {
			t.Fatal(err)
		}
		if len(first) != count || !reflect.DeepEqual(first, second) {
			t.Fatal("порушено довжину або відтворюваність")
		}
		frequencies := map[int]int{}
		categories := [4]bool{}
		for _, value := range first {
			if value < 0 || value > 200000000 {
				t.Fatalf("поза діапазоном: %d", value)
			}
			frequencies[value]++
			switch {
			case value <= 1000000:
				categories[0] = true
			case value <= 10000000:
				categories[1] = true
			case value <= 50000000:
				categories[2] = true
			default:
				categories[3] = true
			}
		}
		if frequencies[5000000] < 2 || categories != [4]bool{true, true, true, true} {
			t.Fatal("немає повтору або категорії")
		}
		other, err := releaseassets.GenerateWithSeed(count, "gohugoio/hugo", 42)
		if err != nil {
			t.Fatal(err)
		}
		if reflect.DeepEqual(first, other) {
			t.Fatal("параметр не вплинув на набір")
		}
	}
	for _, count := range []int{-1, 0, 4, 101} {
		if _, err := releaseassets.Generate(count, "cli/cli"); err == nil {
			t.Fatalf("прийнято count=%d", count)
		}
	}
	for _, subject := range []string{"", " \t "} {
		if _, err := releaseassets.Generate(5, subject); err == nil {
			t.Fatal("прийнято порожній параметр")
		}
	}
	values, err := releaseassets.Generate(5, "cli/cli")
	if err != nil || len(values) != 5 {
		t.Fatal("звичайна генерація не працює")
	}
}
