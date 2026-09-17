// Package launchintervals генерує інтервали між космічними запусками у годинах.
package launchintervals

import "github.com/compilyator/VNTU_GoLang/labs/Lab2Data/internal/generator"

func Generate(count int, provider string) ([]int, error) {
	return GenerateWithSeed(count, provider, generator.Seed())
}
func GenerateWithSeed(count int, provider string, seed int64) ([]int, error) {
	return generator.Ints(count, provider, seed, 1, 720, 1, []int{6, 36, 168, 540, 36})
}
