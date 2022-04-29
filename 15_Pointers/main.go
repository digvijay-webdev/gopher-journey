package main

import "fmt"

func main() {
	// getting the memory address of a variable
	var age int = 16
	memory_age := &age // '&' is used to get memory address
	fmt.Println("Memory Address:", memory_age)

	// getting the value that is stored inside the memory address
	fmt.Println("Stored Value:", *memory_age)

	// updating value of memory address
	*memory_age = 18 //'*' is used to get value stored memory address
	fmt.Println("Updated Value:", *memory_age)
}
