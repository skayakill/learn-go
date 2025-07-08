package main

import "fmt"

func main() {
	// for i := 1;i <= 100; i++ {
	// 	fmt.Println("Angka" , i)
	// }
	for i:= 0;i < 5; i++ {
		for j := i;j < 5; j++ {
			fmt.Print(j,"")
		}
		fmt.Println()
	}
}