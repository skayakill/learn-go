package main

import ("fmt"
		"os"
)

func main() {
	defer fmt.Println("halo")
	fmt.Println("selamat datang")
	os.Exit(1)
	fmt.Println("Terima kasih")
}