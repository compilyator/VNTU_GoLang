// Package biodiversity генерує річні кількості спостережень біорізноманіття.
package biodiversity

import "github.com/compilyator/VNTU_GoLang/labs/Lab2Data/internal/generator"

func Generate(count int, taxon string) ([]int, error) {
	return GenerateWithSeed(count, taxon, generator.Seed())
}
func GenerateWithSeed(count int, taxon string, seed int64) ([]int, error) {
	return generator.Ints(count, taxon, seed, 0, 100000, 10, []int{40, 820, 9400, 68400, 820})
}
