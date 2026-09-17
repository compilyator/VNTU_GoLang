// Package precipitation генерує добові суми опадів у міліметрах.
package precipitation

import "github.com/compilyator/VNTU_GoLang/labs/Lab2Data/internal/generator"

func Generate(count int, city string) ([]float64, error) {
	return GenerateWithSeed(count, city, generator.Seed())
}
func GenerateWithSeed(count int, city string, seed int64) ([]float64, error) {
	return generator.Floats(count, city, seed, 0, 150, 0.1, []float64{0, 3.2, 18.7, 72.4, 0})
}
