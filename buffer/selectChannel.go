package main

import (
	"runtime"
	"fmt"
)

func getAverage(numbers []int, ch chan float64) {
	var sum = 0
	for _, e := range numbers {
		sum += e
	}
	ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers[]int, ch chan int) {
	var max = numbers[0]
	for _, e := range numbers {
		if max < e {
			max = e
		}
	}
	ch <- max
}

func main () {
	runtime.GOMAXPROCS(2)
	var numbers = []int{3, 4, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
	fmt.Println("numbers :", numbers)
	var ch1 = make (chan float64)
	go getAverage(numbers, ch1)
	var ch2 = make(chan int)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Println("avg \t: %.2f \n", avg)
		case max := <-ch2:
			fmt.Println("Max \t : %d \n", max)
		}
	}
}