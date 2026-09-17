// Package musicplays генерує кількості прослуховувань виконавців за вибраний період.
package musicplays

import "github.com/compilyator/VNTU_GoLang/labs/Lab2Data/internal/generator"

func Generate(count int, artist string) ([]int, error) {
	return GenerateWithSeed(count, artist, generator.Seed())
}
func GenerateWithSeed(count int, artist string, seed int64) ([]int, error) {
	return generator.Ints(count, artist, seed, 0, 10000000, 100, []int{1200, 85000, 740000, 6200000, 85000})
}
