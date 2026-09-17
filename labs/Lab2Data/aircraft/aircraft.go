// Package aircraft генерує барометричні висоти літаків у метрах.
package aircraft

import "github.com/compilyator/VNTU_GoLang/labs/Lab2Data/internal/generator"

func Generate(count int, airport string) ([]float64, error) {
	return GenerateWithSeed(count, airport, generator.Seed())
}
func GenerateWithSeed(count int, airport string, seed int64) ([]float64, error) {
	return generator.Floats(count, airport, seed, 0, 13000, 10, []float64{650, 3200, 7600, 11200, 3200})
}
