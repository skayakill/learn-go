package main

import (
	"fmt"
	"strings"
	"errors"
)


func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("kata tidak boleh kosong")
	}
	return true, nil
 }
func main() {
	var name string
	fmt.Println("type your name: ")
	fmt.Scanln(&name)
if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		panic(err.Error())
		fmt.Println("end")
	}	
}