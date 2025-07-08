package main

import (
	"strings"
	"fmt"
)

func main() {
	var names = []string{"My", "skill"}
	printMessage("halo", names)
}
func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString	)
}