package main

type Number interface {
	~int | ~int64 | ~float64
}

func Contains[T comparable](values []T, target T) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

func (s Set[T]) Contains(value T) bool {
	_, ok := s[value]
	return ok
}

func (s Set[T]) Len() int {
	return len(s)
}
