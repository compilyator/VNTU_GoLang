// Package goldprice генерує ціни грама золота в польських злотих.
package goldprice

import "github.com/compilyator/VNTU_GoLang/labs/Lab2Data/internal/generator"

func Generate(count int, year string) ([]float64, error) {
	return GenerateWithSeed(count, year, generator.Seed())
}
func GenerateWithSeed(count int, year string, seed int64) ([]float64, error) {
	return generator.Floats(count, year, seed, 100, 600, 0.01, []float64{165.40, 245.75, 365.20, 515.85, 245.75})
}
