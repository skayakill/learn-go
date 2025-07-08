package main

import "fmt"

func main() {
	// var bilanganBulat uint8= 20
	// var bilanganDecimal = 2.1
	// var varBool = true
	var numberA int = 4
	var numberB *int = &numberA
	*numberB = 8

	fmt.Println("Hello",  numberA, "!\n")
}