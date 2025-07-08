package main

import "fmt"

func main() {
	var firstName string = "Ferry"
	const lastName = "Sabali"

	firstName = "Budi"
	lastName = "Verdy"

	fmt.Print("hello", firstName, lastName, "!\n")


}