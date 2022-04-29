package main

import (
	"fmt"
	"sort"
)

func main() {
	// sorting slice of integers
	slice_1 := []int{3, 4, 6, 1, 2}
	fmt.Println(slice_1)

	sort.Ints(slice_1)
	fmt.Println(slice_1)

	// sorting slice of float
	var slice_2 = []float64{1, 2, 1.2, 0.5, 0.9, 5.4}
	sort.Float64s(slice_2)
	fmt.Println(slice_2)

	// sorting slice of strings
	slice_3 := []string{"C", "A", "b", "B", "D", "d", "k", "p", "G", "h"}
	sort.Strings(slice_3)
	fmt.Println(slice_3)

	// sorting an array
	arr_1 := [3]string{"C", "B", "a"}
	arr_2 := [5]int{4, 3, 6, 1, 2}
	sort.Strings(arr_1[:])
	sort.Ints(arr_2[:])
	fmt.Println(arr_1, arr_2)

	// sorting a struct by one property
	// creating slice of structs
	family := []struct {
		name string
		age  int
	}{
		{"Adam", 16},
		{"John", 25},
		{"Gordan", 36},
		{"Piana", 22},
		{"Morgan", 12},
	}
	fmt.Println(family)

	// works just like JavaScript's sort method
	// for using sort.SliceStable() value we're iterating through needs to be a slice else it will panic
	sort.SliceStable(family, func(a, b int) bool {
		return family[a].age < family[b].age
	})
	fmt.Println(family)
	fmt.Println(len(family))
}
