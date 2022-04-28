package main

import "fmt"

/*
Slices are similar to arrays, but are more powerful and flexible.
Like arrays, slices are also used to store multiple values of the same type in a single variable.
However, unlike arrays, the length of a slice can grow and shrink.

> Using the []datatype{values} format
> Using the make() function
*/

func main() {
	// using []datatype{values} format
	var s_one = []int{1, 2, 3} // static declaration
	fmt.Println(s_one)

	// using make() function
	s_two := make([]int, 1)
	s_two[0] = 4
	fmt.Println(s_two)

	// adding items to slice using append function
	s_two = append(s_two, 5)
	fmt.Println(s_two)

	// making one slice from two
	// '...' at the end is called spread operator (same as javascript)
	newSlice := append(s_one, s_two...)
	fmt.Println(newSlice)
}
