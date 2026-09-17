// Package carbonintensity генерує вуглецеву інтенсивність електроенергії у грамах CO2 на кіловат-годину.
package carbonintensity

import "github.com/compilyator/VNTU_GoLang/labs/Lab2Data/internal/generator"

func Generate(count int, region string) ([]int, error) {
	return GenerateWithSeed(count, region, generator.Seed())
}
func GenerateWithSeed(count int, region string, seed int64) ([]int, error) {
	return generator.Ints(count, region, seed, 0, 500, 1, []int{45, 135, 265, 420, 135})
}
