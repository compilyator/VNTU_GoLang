// Package cryptomarket генерує добові обсяги торгів криптовалютою в мільйонах доларів США.
package cryptomarket

import "github.com/compilyator/VNTU_GoLang/labs/Lab2Data/internal/generator"

func Generate(count int, cryptocurrency string) ([]float64, error) {
	return GenerateWithSeed(count, cryptocurrency, generator.Seed())
}
func GenerateWithSeed(count int, cryptocurrency string, seed int64) ([]float64, error) {
	return generator.Floats(count, cryptocurrency, seed, 1, 100000, 0.1, []float64{120.5, 2450.8, 18200.4, 76500.9, 2450.8})
}
