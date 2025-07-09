package main

import "fmt"

func main() {
	var avg = calculated ( 2, 4, 3, 5, 3 ,3, 5, 5, 3)
	var msg = fmt.Sprintf("rata-rata : %.2f", avg)
	fmt.Println(msg)
}

func calculated(numbers ... int)  float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	var avg = float64(total) / float64(len(numbers))
	
	return avg
}