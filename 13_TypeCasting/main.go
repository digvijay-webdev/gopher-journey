package main

import (
	"fmt"
	"strconv"
)

func main() {
	// converting integers to string
	var integer int = 8
	fmt.Printf("%t", strconv.Itoa(integer))

	// converting floats to string
	val_1 := fmt.Sprint(10.09)
	fmt.Println(val_1)

	// converting string to integer and float
	var str_1 string = "12"
	var str_2 string = "12.08"
	fmt.Printf("%t: %t\n", str_1, str_2)
	newStr_1, err := strconv.ParseInt(str_1, 10, 32)
	newStr_2, err := strconv.ParseFloat(str_2, 64)
	fmt.Printf("%t: %t\n", newStr_1, newStr_2)
	if err == nil {
		fmt.Println("NO Errors..")
	}

	// converting string to boolean and other way around
	var bool_1 string = "true"
	var bool_2 bool = false
	fmt.Printf("%t: %t \n", bool_1, bool_2)

	newBool_1, err := strconv.ParseBool(bool_1)
	newBool_2 := fmt.Sprint(bool_2)
	fmt.Printf("%t\n", newBool_1, newBool_2)
	fmt.Println(err)
}
