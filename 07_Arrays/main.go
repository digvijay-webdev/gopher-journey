package main

import "fmt"

/* Arrays are used to store multiple values of the same type in a single variable,
instead of declaring separate variables for each value.

In Go there are two ways of declaring an array:
1) with specified length
2) without
*/
func main() {
	// format: <'var' keyword> <'name'> = [length]datatype{values}
	var numbers = [3]int{1, 2, 3}
	fmt.Println(numbers)

	var numbersTwo = []int{}
	fmt.Println(numbersTwo)
	numbersTwo = []int{4, 5}
	fmt.Println(numbersTwo)

	// accessing elements of an array
	fmt.Println(numbers[0])
	// changing elements by index
	numbers[0] = 0
	fmt.Println(numbers)

	// we can also assign to specific values and assign to other later.
	// The default value of the skipped indexes will be zero (0)
	numbersThree := []int{0: 10, 2: 30}
	fmt.Println(numbersThree)
	fmt.Println(numbersThree[1]) // in this case index [1] is not assigned

	// getting length of an array using len() function
	fmt.Println(len(numbers))

	// looping through an array
	names := [4]string{}
	names[0] = "John"
	names[1] = "Jane"
	names[2] = "Jimmy"
	names[3] = "Joma"

	fmt.Println(names)

	for nameIndex := 0; nameIndex <= len(names)-1; nameIndex += 1 {
		fmt.Println("Logged: " + names[nameIndex])
	}
}
