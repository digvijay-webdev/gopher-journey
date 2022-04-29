package main

import "fmt"

/*
A struct (short for structure) is used to create a collection of members of different data types,
into a single variable.
While arrays are used to store multiple values of the same data type into a single variable,
structs are used to store multiple values of different data types into a single variable.
A struct can be useful for grouping data together to create records.
Store data in key, value pairs but outputs value only.
*/

// To declare a structure in Go, use the type and struct keywords
type Person struct {
	firstName string
	lastName  string
	age       int
	height    float32
	isAlive   bool
	contact   Contact
}

// nesting struct
type Contact struct {
	email  string
	mobile string
}

func getDetails(person Person) {
	fmt.Println("Name: ", person.firstName)
	fmt.Println("Surname: ", person.lastName)
	fmt.Println("Age: ", person.age)
	fmt.Println("Height: ", person.height)
	fmt.Println("Is Alive: ", person.isAlive)
}

func main() {
	// providing value to Person struct method 1
	var personOne Person
	personOne.firstName = "John"
	personOne.lastName = "Doe"
	personOne.age = 27
	personOne.height = 1.62
	personOne.isAlive = true

	fmt.Println(personOne)
	fmt.Println("Name: " + personOne.firstName) // getting individual property

	// by providing struct as an argument in function
	getDetails(personOne)

	// providing value to Person struct method 2
	// example of nested struct
	var personTwo Person = Person{
		firstName: "Jane",
		lastName:  "Doe",
		age:       26,
		height:    1.43,
		isAlive:   false,
		contact: Contact{
			email:  "jane@doe.com",
			mobile: "0912345678",
		},
	}
	fmt.Println(personTwo)
}
