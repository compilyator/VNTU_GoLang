// Package waves генерує висоти морських хвиль у метрах.
package waves

import "github.com/compilyator/VNTU_GoLang/labs/Lab2Data/internal/generator"

func Generate(count int, seaArea string) ([]float64, error) {
	return GenerateWithSeed(count, seaArea, generator.Seed())
}
func GenerateWithSeed(count int, seaArea string, seed int64) ([]float64, error) {
	return generator.Floats(count, seaArea, seed, 0, 12, 0.1, []float64{0.2, 1.1, 3.4, 7.8, 1.1})
}
