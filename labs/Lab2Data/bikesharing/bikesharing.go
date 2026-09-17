// Package bikesharing генерує кількості доступних велосипедів на станціях прокату.
package bikesharing

import "github.com/compilyator/VNTU_GoLang/labs/Lab2Data/internal/generator"

func Generate(count int, city string) ([]int, error) {
	return GenerateWithSeed(count, city, generator.Seed())
}
func GenerateWithSeed(count int, city string, seed int64) ([]int, error) {
	return generator.Ints(count, city, seed, 0, 80, 1, []int{0, 4, 18, 52, 4})
}
