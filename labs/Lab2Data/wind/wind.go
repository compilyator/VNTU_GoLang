// Package wind генерує максимальні добові швидкості вітру в кілометрах за годину.
package wind

import "github.com/compilyator/VNTU_GoLang/labs/Lab2Data/internal/generator"

func Generate(count int, city string) ([]float64, error) {
	return GenerateWithSeed(count, city, generator.Seed())
}
func GenerateWithSeed(count int, city string, seed int64) ([]float64, error) {
	return generator.Floats(count, city, seed, 0, 150, 0.1, []float64{2.4, 18.6, 46.2, 96.8, 18.6})
}
