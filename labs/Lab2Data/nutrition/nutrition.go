// Package nutrition генерує вміст цукру у грамах на 100 грамів продукту.
package nutrition

import "github.com/compilyator/VNTU_GoLang/labs/Lab2Data/internal/generator"

func Generate(count int, category string) ([]float64, error) {
	return GenerateWithSeed(count, category, generator.Seed())
}
func GenerateWithSeed(count int, category string, seed int64) ([]float64, error) {
	return generator.Floats(count, category, seed, 0, 80, 0.1, []float64{2.5, 11.8, 28.4, 61.7, 11.8})
}
