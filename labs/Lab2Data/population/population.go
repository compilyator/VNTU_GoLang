// Package population генерує річні значення населення в тисячах осіб.
package population

import "github.com/compilyator/VNTU_GoLang/labs/Lab2Data/internal/generator"

func Generate(count int, country string) ([]int, error) {
	return GenerateWithSeed(count, country, generator.Seed())
}
func GenerateWithSeed(count int, country string, seed int64) ([]int, error) {
	return generator.Ints(count, country, seed, 500, 100000, 100, []int{3500, 18000, 42000, 76000, 42000})
}
