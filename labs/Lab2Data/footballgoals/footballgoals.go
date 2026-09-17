// Package footballgoals генерує загальні кількості голів у футбольних матчах.
package footballgoals

import "github.com/compilyator/VNTU_GoLang/labs/Lab2Data/internal/generator"

func Generate(count int, competition string) ([]int, error) {
	return GenerateWithSeed(count, competition, generator.Seed())
}
func GenerateWithSeed(count int, competition string, seed int64) ([]int, error) {
	return generator.Ints(count, competition, seed, 0, 12, 1, []int{0, 2, 4, 7, 2})
}
