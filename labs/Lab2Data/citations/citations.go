// Package citations генерує кількості цитувань наукових публікацій.
package citations

import "github.com/compilyator/VNTU_GoLang/labs/Lab2Data/internal/generator"

func Generate(count int, topic string) ([]int, error) {
	return GenerateWithSeed(count, topic, generator.Seed())
}
func GenerateWithSeed(count int, topic string, seed int64) ([]int, error) {
	return generator.Ints(count, topic, seed, 0, 5000, 1, []int{0, 8, 72, 840, 8})
}
