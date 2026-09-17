// Package cryptocandles генерує добові ціни закриття криптовалюти в доларах США.
package cryptocandles

import "github.com/compilyator/VNTU_GoLang/labs/Lab2Data/internal/generator"

func Generate(count int, pair string) ([]float64, error) {
	return GenerateWithSeed(count, pair, generator.Seed())
}
func GenerateWithSeed(count int, pair string, seed int64) ([]float64, error) {
	return generator.Floats(count, pair, seed, 10000, 150000, 0.01, []float64{28450.25, 52400.80, 78250.45, 126800.10, 52400.80})
}
