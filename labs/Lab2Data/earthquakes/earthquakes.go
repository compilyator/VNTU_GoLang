// Package earthquakes генерує магнітуди землетрусів.
package earthquakes

import "github.com/compilyator/VNTU_GoLang/labs/Lab2Data/internal/generator"

func Generate(count int, region string) ([]float64, error) {
	return GenerateWithSeed(count, region, generator.Seed())
}
func GenerateWithSeed(count int, region string, seed int64) ([]float64, error) {
	return generator.Floats(count, region, seed, 1, 9.5, 0.1, []float64{2.6, 4.7, 6.3, 8.1, 4.7})
}
