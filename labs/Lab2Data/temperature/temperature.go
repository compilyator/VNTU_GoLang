// Package temperature генерує середні добові температури повітря у градусах Цельсія.
package temperature

import "github.com/compilyator/VNTU_GoLang/labs/Lab2Data/internal/generator"

func Generate(count int, city string) ([]float64, error) {
	return GenerateWithSeed(count, city, generator.Seed())
}
func GenerateWithSeed(count int, city string, seed int64) ([]float64, error) {
	return generator.Floats(count, city, seed, -30, 45, 0.1, []float64{-8.5, 8.2, 21.4, 34.6, 21.4})
}
