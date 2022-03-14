package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	min, max := 100, 1000
	arr := [10]int{}

	for i := range arr {
		arr[i] = rand.Intn(max-min) + min
	}
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func(ch chan int) {
		for _, v := range arr {
			ch <- v
		}
		close(ch1)
	}(ch1)

	go func(ch1, ch2 chan int) {
		for i := range ch1 {
			ch2 <- i * i
		}
		close(ch2)
	}(ch1, ch2)

	for i := range ch2 {
		fmt.Printf("Значение массива: %v\n", i)
	}

}
