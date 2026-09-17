// Package pageviews генерує добові кількості переглядів статті Вікіпедії.
package pageviews

import "github.com/compilyator/VNTU_GoLang/labs/Lab2Data/internal/generator"

func Generate(count int, article string) ([]int, error) {
	return GenerateWithSeed(count, article, generator.Seed())
}
func GenerateWithSeed(count int, article string, seed int64) ([]int, error) {
	return generator.Ints(count, article, seed, 0, 1000000, 100, []int{800, 12500, 74000, 420000, 12500})
}
