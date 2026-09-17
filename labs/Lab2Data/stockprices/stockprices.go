// Package stockprices генерує добові ціни закриття акцій у доларах США.
package stockprices

import "github.com/compilyator/VNTU_GoLang/labs/Lab2Data/internal/generator"

func Generate(count int, ticker string) ([]float64, error) {
	return GenerateWithSeed(count, ticker, generator.Seed())
}
func GenerateWithSeed(count int, ticker string, seed int64) ([]float64, error) {
	return generator.Floats(count, ticker, seed, 1, 1000, 0.01, []float64{24.80, 118.35, 346.90, 812.45, 118.35})
}
