package main

import (
	"errors"
	"fmt"
	"os"
)

/*
Go doesn't have exceptions, thus it does not support try and catch syntax. But for error handling we have
multiple return values and panic() function.
*/

func main() {
	result, err := os.ReadFile("./test1.txt")
	handleError(err, "Failed to create file..")
	fmt.Println("File created..", result)
}

// Panic is a built-in function that stops the normal execution flow.
// The deferred functions are still run as usual.
// Panic function will exit the app/process.
// errors.New() method is used to create custom error message upon calling panic function.
func handleError(err error, possibleErrMessage string) {
	if err != nil {
		fmt.Println(err)
		panic(errors.New(possibleErrMessage))
	}
}
