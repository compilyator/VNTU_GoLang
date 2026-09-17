// Package questions генерує рейтинги запитань Stack Exchange.
package questions

import "github.com/compilyator/VNTU_GoLang/labs/Lab2Data/internal/generator"

func Generate(count int, tag string) ([]int, error) {
	return GenerateWithSeed(count, tag, generator.Seed())
}
func GenerateWithSeed(count int, tag string, seed int64) ([]int, error) {
	return generator.Ints(count, tag, seed, -20, 1000, 1, []int{-3, 0, 18, 240, 0})
}
