package main

import "fmt"

func main() {
	var point = 6
	switch point {
	case 7:
		fmt.Printf("perfect!")
	case 6:
		fmt.Printf("excellent")
	case 5:
		fmt.Printf("Good")
	default:
	fmt.Printf("Not Bad")
	}
}