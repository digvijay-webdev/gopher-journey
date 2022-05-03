package main

import (
	"encoding/json"
	"fmt"
)

// struct declaration
type Units struct {
	H string
	W string
}

type Body struct {
	Height float32
	Weight float32
	Units  Units
}

type Person struct {
	Name   string
	IsMale bool
	Age    int
	Body   Body
}

func main() {
	// parsing struct to JSON
	var JSON_OBJ string = structToJSON()
	fmt.Println(JSON_OBJ)

	// parsing JSON to objects
	var object Person
	err := json.Unmarshal([]byte(JSON_OBJ), &object)
	if err != nil {
		panic(err)
	}
	fmt.Println(object)
}

// parsing struct to json
// when parsing to json every property on struct must be capitalized
// When the first letter is capitalised, the identifier is public to any piece of code that you want to use.
// When the first letter is lowercase, the identifier is private and could only be accessed within the package it was declared.
// we can similarly use json.Marshal() to parse maps as well.
func structToJSON() string {
	personOne := Person{
		Name:   "John Doe",
		IsMale: true,
		Age:    32,
		Body: Body{
			Height: 172.5,
			Weight: 95.89,
			Units: Units{
				H: "cm",
				W: "kg",
			},
		},
	}

	// simple log
	fmt.Println(personOne)

	// parsing into json
	bytes, err := json.MarshalIndent(personOne, "", "\t") // for printing JSON with proper indentation
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
