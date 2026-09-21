package main

import "testing"

type Score int

func TestContains(t *testing.T) {
	if !Contains([]int{1, 2, 3}, 2) {
		t.Fatal("Contains() = false, want true")
	}
	if Contains([]string{"go"}, "web") {
		t.Fatal("Contains() = true, want false")
	}
}

func TestMinSupportsDefinedType(t *testing.T) {
	if got, want := Min(Score(8), Score(3)), Score(3); got != want {
		t.Fatalf("Min() = %v, want %v", got, want)
	}
}

func TestSetKeepsUniqueValues(t *testing.T) {
	set := Set[string]{}
	set.Add("go")
	set.Add("go")
	set.Add("web")
	if got, want := set.Len(), 2; got != want {
		t.Fatalf("Len() = %d, want %d", got, want)
	}
}
