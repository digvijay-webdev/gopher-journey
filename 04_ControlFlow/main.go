package main

import "fmt"

var age int
var hasLicense bool = true

func main() {
	fmt.Println("Please enter your age:")
	fmt.Scan(&age)

	if age >= 18 && age < 100 {
		if hasLicense {
			fmt.Println("Can drive")
		} else {
			fmt.Println("Cannot drive")
		}
	} else if age <= 0 || age > 100 {
		fmt.Println("invalid input")
	} else {
		fmt.Println("cannot drive")
	}
}
