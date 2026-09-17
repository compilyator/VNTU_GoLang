package lab2data_test

import (
	"reflect"
	"testing"

	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/aircraft"
	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/airquality"
	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/asteroids"
	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/bikesharing"
	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/biodiversity"
	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/books"
	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/carbonintensity"
	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/citations"
	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/cryptocandles"
	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/cryptomarket"
	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/earthquakes"
	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/exchange"
	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/f1laps"
	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/footballgoals"
	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/goldprice"
	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/launchintervals"
	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/musicplays"
	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/naturalevents"
	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/nburate"
	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/nutrition"
	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/pageviews"
	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/population"
	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/precipitation"
	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/questions"
	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/socialindicator"
	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/stockprices"
	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/temperature"
	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/waterflow"
	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/waves"
	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/wind"
)

func TestFloatGeneratorsAreDeterministic(t *testing.T) {
	tests := map[string]func(int, string, int64) ([]float64, error){
		"aircraft": aircraft.GenerateWithSeed, "airquality": airquality.GenerateWithSeed,
		"asteroids": asteroids.GenerateWithSeed, "cryptocandles": cryptocandles.GenerateWithSeed,
		"cryptomarket": cryptomarket.GenerateWithSeed, "earthquakes": earthquakes.GenerateWithSeed,
		"exchange": exchange.GenerateWithSeed, "goldprice": goldprice.GenerateWithSeed,
		"nburate": nburate.GenerateWithSeed, "nutrition": nutrition.GenerateWithSeed,
		"precipitation": precipitation.GenerateWithSeed, "socialindicator": socialindicator.GenerateWithSeed,
		"stockprices": stockprices.GenerateWithSeed, "temperature": temperature.GenerateWithSeed,
		"waterflow": waterflow.GenerateWithSeed, "waves": waves.GenerateWithSeed, "wind": wind.GenerateWithSeed,
	}
	for name, generate := range tests {
		t.Run(name, func(t *testing.T) {
			first, err := generate(20, "основний об’єкт", 42)
			if err != nil {
				t.Fatal(err)
			}
			second, err := generate(20, "основний об’єкт", 42)
			if err != nil {
				t.Fatal(err)
			}
			if len(first) != 20 || !reflect.DeepEqual(first, second) {
				t.Fatalf("контракт порушено: %v / %v", first, second)
			}
			if !hasFloatDuplicate(first) {
				t.Fatalf("генератор не створив повтор: %v", first)
			}
			other, err := generate(20, "інший об’єкт", 42)
			if err != nil {
				t.Fatal(err)
			}
			if reflect.DeepEqual(first, other) {
				t.Fatal("предметний параметр не вплинув на набір")
			}
			if _, err := generate(20, "  ", 42); err == nil {
				t.Fatal("очікувалася помилка для порожнього предметного параметра")
			}
		})
	}
}

func TestIntegerGeneratorsAreDeterministic(t *testing.T) {
	tests := map[string]func(int, string, int64) ([]int, error){
		"bikesharing": bikesharing.GenerateWithSeed, "biodiversity": biodiversity.GenerateWithSeed,
		"books": books.GenerateWithSeed, "carbonintensity": carbonintensity.GenerateWithSeed,
		"citations": citations.GenerateWithSeed, "f1laps": f1laps.GenerateWithSeed,
		"footballgoals": footballgoals.GenerateWithSeed, "launchintervals": launchintervals.GenerateWithSeed,
		"musicplays": musicplays.GenerateWithSeed, "naturalevents": naturalevents.GenerateWithSeed,
		"pageviews": pageviews.GenerateWithSeed, "population": population.GenerateWithSeed,
		"questions": questions.GenerateWithSeed,
	}
	for name, generate := range tests {
		t.Run(name, func(t *testing.T) {
			first, err := generate(20, "основний об’єкт", 42)
			if err != nil {
				t.Fatal(err)
			}
			second, err := generate(20, "основний об’єкт", 42)
			if err != nil {
				t.Fatal(err)
			}
			if len(first) != 20 || !reflect.DeepEqual(first, second) {
				t.Fatalf("контракт порушено: %v / %v", first, second)
			}
			if !hasIntDuplicate(first) {
				t.Fatalf("генератор не створив повтор: %v", first)
			}
			other, err := generate(20, "інший об’єкт", 42)
			if err != nil {
				t.Fatal(err)
			}
			if reflect.DeepEqual(first, other) {
				t.Fatal("предметний параметр не вплинув на набір")
			}
			if _, err := generate(20, "\t", 42); err == nil {
				t.Fatal("очікувалася помилка для порожнього предметного параметра")
			}
		})
	}
}

func hasFloatDuplicate(values []float64) bool {
	seen := make(map[float64]bool, len(values))
	for _, value := range values {
		if seen[value] {
			return true
		}
		seen[value] = true
	}
	return false
}

func hasIntDuplicate(values []int) bool {
	seen := make(map[int]bool, len(values))
	for _, value := range values {
		if seen[value] {
			return true
		}
		seen[value] = true
	}
	return false
}

func TestGeneratorsRejectCountOutsideRange(t *testing.T) {
	if _, err := temperature.GenerateWithSeed(4, "Вінниця", 1); err == nil {
		t.Fatal("очікувалася помилка для count=4")
	}
	if _, err := footballgoals.GenerateWithSeed(101, "Premier League", 1); err == nil {
		t.Fatal("очікувалася помилка для count=101")
	}
}

func TestSubjectNormalization(t *testing.T) {
	first, err := temperature.GenerateWithSeed(20, "  ВІННИЦЯ ", 42)
	if err != nil {
		t.Fatal(err)
	}
	second, err := temperature.GenerateWithSeed(20, "вінниця", 42)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(first, second) {
		t.Fatal("регістр і зовнішні пробіли не мають змінювати набір")
	}
}

func TestGenerateRejectsBlankSubject(t *testing.T) {
	if _, err := temperature.Generate(5, " \t "); err == nil {
		t.Fatal("очікувалася помилка для порожнього предметного параметра")
	}
}
