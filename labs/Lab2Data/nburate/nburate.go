// Package nburate генерує офіційні курси гривні.
package nburate

import "github.com/compilyator/VNTU_GoLang/labs/Lab2Data/internal/generator"

func Generate(count int, currency string) ([]float64, error) {
	return GenerateWithSeed(count, currency, generator.Seed())
}
func GenerateWithSeed(count int, currency string, seed int64) ([]float64, error) {
	return generator.Floats(count, currency, seed, 20, 80, 0.01, []float64{27.42, 38.15, 46.73, 67.28, 38.15})
}
