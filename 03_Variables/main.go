package main

import "fmt"

var glob1, glob2 = "John", "Jane"

func main() {
	// Declaration: <'var' keyword> <variable name> <datatype> <'=' assignment operator> <'value'>
	// integers
	var age int = 21
	fmt.Println(age)

	// we can declare the variable and assign the value later
	var rollNumber int
	fmt.Println("Array Starts from:")
	fmt.Println(rollNumber)
	rollNumber = 28
	fmt.Println("My roll number is:")
	fmt.Println(rollNumber)

	// float32- stores floating point numbers, with decimals
	var Pi float32 = 3.14
	fmt.Println("Pi:")
	fmt.Println(Pi)

	// string
	var name string = "John Doe"
	fmt.Println("Hello my name is " + name)

	// bool || boolean
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)

	// we can dynamically assign the value to variable like javascript
	var tempOne = "test 01"
	fmt.Println(tempOne)                                              // string
	tempOne = "Can change the value but has to be the same data type" // read <=
	fmt.Println(tempOne)

	// can use := operator to achive the same result
	tempTwo := "test 02"
	fmt.Println(tempTwo)
	tempTwo = "Test 02"
	fmt.Println(tempTwo)

	/*
		Note:
		':=' this operator won't work outside the functions. use 'var' declaration for using it globally.
		Variable declaration must follow the naming rules like any other language.
	*/

	// declaring multiple variables at once
	fmt.Println(glob1, glob2)

	// Constants
	const constInfo = "Constants are similar to variables but they can't change their value after declaration."
	fmt.Println(constInfo)
}
