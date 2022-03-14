package main

import (
	"fmt"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10}

	ch := make(chan int)

	go func(ch chan int, arr *[5]int) {
		for _, v := range *arr {
			ch <- v * v
		}
		close(ch)
	}(ch, &arr)

	fmt.Printf("Значение квадратов массива %v = ", arr)

	for i := range ch {
		fmt.Printf("%v ", i)
	}
	fmt.Printf("\n")
}
