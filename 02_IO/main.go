package main

import "fmt"

func main() {
	// prints on new line
	fmt.Println("Use me for printing on a new line..")

	// prints on the same line
	fmt.Print("Same Line.\nNew line begins.")

	// by default will print first item/character of input
	var name string
	fmt.Scanln(&name)
	fmt.Println(name)

	// format string
	var firstName, age = "John", 40
	fmt.Printf("Hello my name is %v and i am %v years old", firstName, age)
}
