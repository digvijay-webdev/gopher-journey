package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	execute()
}

func execute() {
	fmt.Println("Working with files..")

	// creating file
	create()

	// renaming or moving a file
	rename()

	// coping file
	copy()

	// delete file
	delete()

	// create and write file
	createAndWrite()

	// reading a file
	read()
}

// create will not create folders by itself
// you either have to create them manually or program dynamically before performing the file create operation
func create() {
	result, err := os.Create("./files/testFile.html")
	if err != nil {
		fmt.Println("Error: Failed to create file..")
		fmt.Println(err)
		return
	}

	fmt.Println("File Created..", result)
}

// rename method can rename the file but could also be used to move a file to different location
// to move a file with other name, in different location just edit the newPath (2nd param) accordingly
func rename() {
	err := os.Rename("./files/testFile.html", "./files/goTestFile.html")
	if err != nil {
		fmt.Println("Error: Failed to rename file..")
		fmt.Println(err)
		return
	}

	fmt.Println("File Renamed..")
}

// there are many ways of copying files in go io.Copy() is one of them
func copy() {
	// step: 1 - open the file
	originalFile, err := os.Open("./files/goTestFile.html")
	if err != nil {
		fmt.Println("Error: Failed to copy the file..")
		fmt.Println(err)
		return
	}
	// step: 2 - closing the file using defer to delay the close operation
	defer originalFile.Close()

	// step: 3 - create new file
	newFile, err := os.Create("./files/copy.txt")
	if err != nil {
		fmt.Println("Error: Failed to create new file..")
		fmt.Println(err)
		return
	}
	defer newFile.Close()

	// step: 4 - write data in bytes
	dataInBytes, err := io.Copy(newFile, originalFile)
	if err != nil {
		fmt.Println("Error: Failed to write bytes to copy..")
		fmt.Println(err)
		return
	}
	fmt.Println("Bytes written..", dataInBytes)
	fmt.Println("File copied..", originalFile)
}

// os.Remove() doesn't return any result the file will be deleted if there is not error
func delete() {
	err := os.Remove("./files/goTestFile.html")
	if err != nil {
		fmt.Println("Error: Failed to delete file..")
		fmt.Println(err)
		return
	}
	fmt.Println("File deleted..")
}

// get/convert data to 'bytes' before writing it to file
// only returns error object
// first param is path and name of the file
// second param is data in byte format
// third and last param is permission - 0644
// if the file already exist then os.WriteFile() will replace the old one
func createAndWrite() {
	data := []byte("console.log('logged-in test..')")
	err := os.WriteFile("./files/index.js", data, 0644)
	if err != nil {
		fmt.Println("Error: Failed to create and write file..")
		fmt.Println(err)
		return
	}
}

// os.ReadFile() is used to read file's content
// above function takes one param which is the path to the file
// it returns error and data in the buffer format
// to get the data in readable format use string() function to convert it to string
func read() {
	result, err := os.ReadFile("./files/copy.txt")
	if err != nil {
		fmt.Println("Error: Failed to read file..")
		fmt.Println(err)
		return
	}
	fmt.Println("File's content..")
	fmt.Println(string(result))
}
