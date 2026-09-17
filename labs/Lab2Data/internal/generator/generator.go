package generator

import (
	"errors"
	"hash/fnv"
	"math"
	"math/rand"
	"strings"
	"time"
)

const (
	MinCount = 5
	MaxCount = 100
)

var ErrCountOutOfRange = errors.New("кількість значень має бути від 5 до 100")
var ErrSubjectRequired = errors.New("предметний параметр не задано")

func Floats(count int, subject string, seed int64, minValue, maxValue, step float64, anchors []float64) ([]float64, error) {
	if count < MinCount || count > MaxCount {
		return nil, ErrCountOutOfRange
	}
	seed, err := seedWithSubject(seed, subject)
	if err != nil {
		return nil, err
	}
	if step <= 0 || minValue > maxValue || len(anchors) == 0 {
		return nil, errors.New("неправильна конфігурація генератора")
	}

	rng := rand.New(rand.NewSource(seed))
	values := make([]float64, 0, count)
	for len(values) < count && len(values) < len(anchors) {
		values = append(values, anchors[len(values)])
	}
	steps := int(math.Floor((maxValue-minValue)/step)) + 1
	for len(values) < count {
		value := minValue + float64(rng.Intn(steps))*step
		values = append(values, round(value, step))
	}
	rng.Shuffle(len(values), func(i, j int) { values[i], values[j] = values[j], values[i] })
	return values, nil
}

func Ints(count int, subject string, seed int64, minValue, maxValue, step int, anchors []int) ([]int, error) {
	if count < MinCount || count > MaxCount {
		return nil, ErrCountOutOfRange
	}
	seed, err := seedWithSubject(seed, subject)
	if err != nil {
		return nil, err
	}
	if step <= 0 || minValue > maxValue || len(anchors) == 0 {
		return nil, errors.New("неправильна конфігурація генератора")
	}

	rng := rand.New(rand.NewSource(seed))
	values := make([]int, 0, count)
	for len(values) < count && len(values) < len(anchors) {
		values = append(values, anchors[len(values)])
	}
	steps := (maxValue-minValue)/step + 1
	for len(values) < count {
		values = append(values, minValue+rng.Intn(steps)*step)
	}
	rng.Shuffle(len(values), func(i, j int) { values[i], values[j] = values[j], values[i] })
	return values, nil
}

func Seed() int64 {
	return time.Now().UnixNano()
}

func seedWithSubject(seed int64, subject string) (int64, error) {
	normalized := strings.ToLower(strings.TrimSpace(subject))
	if normalized == "" {
		return 0, ErrSubjectRequired
	}

	hash := fnv.New64a()
	_, _ = hash.Write([]byte(normalized))
	return seed ^ int64(hash.Sum64()), nil
}

func round(value, step float64) float64 {
	digits := 0
	for current := step; current < 1 && digits < 9; current *= 10 {
		digits++
	}
	factor := math.Pow10(digits)
	return math.Round(value*factor) / factor
}
