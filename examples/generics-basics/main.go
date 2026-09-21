package main

import "fmt"

func main() {
	fmt.Println(Contains([]string{"go", "web"}, "go"))
	fmt.Println(Min(7, 2))

	values := Set[int]{}
	values.Add(4)
	values.Add(4)
	values.Add(9)
	fmt.Println(values.Len())
}
