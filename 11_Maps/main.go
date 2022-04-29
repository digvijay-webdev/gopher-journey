package main

import "fmt"

/*
Maps are used to store data values in key:value pairs.
Each element in a map is a key:value pair.
A map is an unordered and changeable collection that does not allow duplicates.
The length of a map is the number of its elements. You can find it using the len() function
The default value of a map is nil.
Maps hold references to an underlying hash table
Go has two ways for creating maps: with make() function and using ':=' operator
Maps are references to hash tables.
If two map variables refer to the same hash table, changing the content of one variable affect the content of the other.
*/

func main() {
	// creating a map using ':=' operator
	mapOne := map[string]string{
		"firstName": "John",
		"lastName":  "Doe",
	}
	fmt.Println(mapOne)
	fmt.Println(mapOne["firstName"], mapOne["lastName"])

	// creating a map using make() function
	mapTwo := make(map[string]string)

	// assigning values to map
	mapTwo["firstName"] = "Jane"
	mapTwo["lastName"] = "Doe"
	fmt.Println(mapTwo)

	// updating or adding to map
	mapTwo["firstName"] = "Jema"
	mapTwo["friend"] = "John"
	fmt.Println(mapTwo)

	// deleting a specific key, value pair from the map
	delete(mapTwo, "friend")
	fmt.Println(mapTwo)

	// checking if the key and value exists
	checkKV()
}

func checkKV() {
	//	score board
	var scores = map[string]int{
		"Ironman":    100,
		"Cap":        66,
		"Hulk":       89,
		"Thor":       92,
		"Vision":     78,
		"Warmachine": 45,
	}

	// check if Vision and Hawkeye exist
	val1, ok1 := scores["Vision"] // checking for both value and is exists 'val1' will hold the value and ok1 will result (bool)
	_, ok2 := scores["Hawkey"]    // checking for result only

	fmt.Println("Checking if specific key, values exists..")
	fmt.Println(val1, ok1) // returns 78 and true
	fmt.Println(ok2)       // returns false

	// looping through a map
	for key, value := range scores {
		fmt.Printf("%v: %v\n", key, value)
	}
}
