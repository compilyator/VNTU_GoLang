// Package waterflow генерує середньодобові витрати річкової води в кубічних метрах за секунду.
package waterflow

import "github.com/compilyator/VNTU_GoLang/labs/Lab2Data/internal/generator"

func Generate(count int, river string) ([]float64, error) {
	return GenerateWithSeed(count, river, generator.Seed())
}
func GenerateWithSeed(count int, river string, seed int64) ([]float64, error) {
	return generator.Floats(count, river, seed, 0, 5000, 0.1, []float64{85.2, 620.4, 1820.7, 4120.3, 620.4})
}
