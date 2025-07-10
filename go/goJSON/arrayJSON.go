package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Name"`
	Age int
}

func main() {
	var jsonString = `[
	{"Name" : "Jhon wick", "Age" : 27}, {"Name": "etham hunt", "Age": 32}
	]`
	var data []User
	var err = json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("user 1:", data[0].FullName)
	fmt.Println("user 2:", data[1].FullName)
}