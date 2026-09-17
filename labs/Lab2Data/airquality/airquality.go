// Package airquality генерує концентрації PM2.5 у мікрограмах на кубічний метр.
package airquality

import "github.com/compilyator/VNTU_GoLang/labs/Lab2Data/internal/generator"

func Generate(count int, city string) ([]float64, error) {
	return GenerateWithSeed(count, city, generator.Seed())
}
func GenerateWithSeed(count int, city string, seed int64) ([]float64, error) {
	return generator.Floats(count, city, seed, 0, 250, 0.1, []float64{7.5, 24.8, 48.3, 121.6, 24.8})
}
