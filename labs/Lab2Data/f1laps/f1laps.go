// Package f1laps генерує час кіл Формули-1 у мілісекундах.
package f1laps

import "github.com/compilyator/VNTU_GoLang/labs/Lab2Data/internal/generator"

func Generate(count int, driver string) ([]int, error) {
	return GenerateWithSeed(count, driver, generator.Seed())
}
func GenerateWithSeed(count int, driver string, seed int64) ([]int, error) {
	return generator.Ints(count, driver, seed, 60000, 180000, 10, []int{74820, 89540, 108760, 151230, 89540})
}
