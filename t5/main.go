package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int
	fmt.Printf("Введите количество секунд: ")
	fmt.Scanf("%d\n", &n)
	ch := make(chan int)
	go func(w chan int) {
		for {
			v := rand.Int()
			fmt.Println("Отправляю значение в канал ", v)
			w <- v
		}
	}(ch)

	go func(r chan int) {
		for i := range r {
			fmt.Println("Ловлю значение из канала ", i)
		}
	}(ch)
	time.Sleep(time.Duration(n) * time.Second)

}
