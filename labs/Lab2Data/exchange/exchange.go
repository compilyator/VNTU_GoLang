// Package exchange генерує курси вибраної валютної пари.
package exchange

import "github.com/compilyator/VNTU_GoLang/labs/Lab2Data/internal/generator"

func Generate(count int, pair string) ([]float64, error) {
	return GenerateWithSeed(count, pair, generator.Seed())
}
func GenerateWithSeed(count int, pair string, seed int64) ([]float64, error) {
	return generator.Floats(count, pair, seed, 0.5, 2, 0.001, []float64{0.82, 0.98, 1.08, 1.34, 0.98})
}
