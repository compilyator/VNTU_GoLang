// Package releaseassets генерує навчальні розміри файлів програмних релізів у байтах.
package releaseassets

import "github.com/compilyator/VNTU_GoLang/labs/Lab2Data/internal/generator"

// Generate повертає синтетичні дані, а не відомості з GitHub.
func Generate(count int, repository string) ([]int, error) {
	return GenerateWithSeed(count, repository, generator.Seed())
}

// GenerateWithSeed відтворює набір за кількістю, репозиторієм і seed.
func GenerateWithSeed(count int, repository string, seed int64) ([]int, error) {
	return generator.Ints(count, repository, seed, 0, 200000000, 500000,
		[]int{500000, 5000000, 25000000, 75000000, 5000000})
}
