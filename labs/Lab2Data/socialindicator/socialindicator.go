// Package socialindicator генерує річні значення очікуваної тривалості життя.
package socialindicator

import "github.com/compilyator/VNTU_GoLang/labs/Lab2Data/internal/generator"

func Generate(count int, country string) ([]float64, error) {
	return GenerateWithSeed(count, country, generator.Seed())
}
func GenerateWithSeed(count int, country string, seed int64) ([]float64, error) {
	return generator.Floats(count, country, seed, 40, 90, 0.1, []float64{52.3, 66.8, 75.4, 84.1, 75.4})
}
