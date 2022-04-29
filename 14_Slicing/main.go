package main

import "fmt"

/*
An array as fixed size, meaning once defined, you cannot change the number of elements it holds.
To overcome this Go provides the slice, which is a dynamically sized view into elements of an array.
A slice is based on an array and is defined by specifying two indices,  a low and high bound, separated by colon.

*/

func main() {
	// creating slice
	s1 := [6]int{1, 2, 3, 4, 5, 6}

	// create a new array with first three indexes from s1
	// last index in this case 3rd will not considered
	s1_half := s1[0:3]
	fmt.Println(s1_half)

	// create a new array with last three indexes from s1
	// by leaving [3:] empty will go untill the last element
	// if [:] there is no index specified will return whole array or slice
	s1_otherHalf := s1[3:]
	fmt.Println(s1_otherHalf)

	// we can also slice the string
	str1 := "SLICING A STRING"
	fmt.Println(str1)

	str2 := str1[:1]
	fmt.Println(str2) // prints first character
}
