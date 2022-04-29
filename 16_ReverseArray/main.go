package main

import (
	"fmt"
)

func main() {

	// reversing an array
	arr_1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr_1)
	arr_1_new := reverseIntArray(arr_1)
	fmt.Println(arr_1_new)

	// reversing string
	str_1 := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fmt.Println(str_1)

	str_1_new := reverseString(str_1)
	fmt.Println(str_1_new)
}

// reverse string
func reverseString(value string) (result string) {
	for _, v := range value {
		result = string(v) + result // string() function converts to UTF-8
	}
	return result
}

// reverse array
func reverseIntArray(array []int) []int {
	reversed_arr_1 := []int{}
	for i := len(array) - 1; i >= 0; i-- {
		reversed_arr_1 = append(reversed_arr_1, array[i])
	}
	return reversed_arr_1
}
