// Package asteroids генерує оцінені максимальні діаметри астероїдів у метрах.
package asteroids

import "github.com/compilyator/VNTU_GoLang/labs/Lab2Data/internal/generator"

func Generate(count int, orbitalClass string) ([]float64, error) {
	return GenerateWithSeed(count, orbitalClass, generator.Seed())
}
func GenerateWithSeed(count int, orbitalClass string, seed int64) ([]float64, error) {
	return generator.Floats(count, orbitalClass, seed, 1, 2000, 0.1, []float64{18.6, 82.4, 315.7, 1240.3, 82.4})
}
