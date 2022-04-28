package main

import "fmt"

func main() {
	fmt.Println("Let's loop..")

	// expressions: initializer, condition, increase/decrese value
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// infinite loop & break
	in := 0
	for true {
		if in == 8 {
			break // exists the loop
		}
		fmt.Print("Loop")
		in++
	}
	fmt.Println()

	// continue statement is used to skip one iteration
	for i := 1; i <= 10; i = i + 1 {
		// skip number 8
		if i == 8 {
			continue // we're skipping iteration when i is 8
		}
		fmt.Println(i)
	}
}
