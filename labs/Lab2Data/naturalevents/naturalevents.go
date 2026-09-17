// Package naturalevents генерує тривалості природних подій у днях.
package naturalevents

import "github.com/compilyator/VNTU_GoLang/labs/Lab2Data/internal/generator"

func Generate(count int, category string) ([]int, error) {
	return GenerateWithSeed(count, category, generator.Seed())
}
func GenerateWithSeed(count int, category string, seed int64) ([]int, error) {
	return generator.Ints(count, category, seed, 1, 365, 1, []int{2, 14, 75, 240, 14})
}
