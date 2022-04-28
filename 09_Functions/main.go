package main

import "fmt"

// declaring function
// format: <'func' keyword> <name(parameter datatype)> <return datatype>
// Go does not support default parameters
func greet(name string) string {
	return "Greetings " + name
}

func main() {
	// calling function
	fmt.Println(greet("Tony"))
	fmt.Println(greet("Charit"))
}
