// Package books генерує медіанні кількості сторінок книг.
package books

import "github.com/compilyator/VNTU_GoLang/labs/Lab2Data/internal/generator"

func Generate(count int, category string) ([]int, error) {
	return GenerateWithSeed(count, category, generator.Seed())
}
func GenerateWithSeed(count int, category string, seed int64) ([]int, error) {
	return generator.Ints(count, category, seed, 20, 2000, 1, []int{84, 245, 620, 1380, 245})
}
