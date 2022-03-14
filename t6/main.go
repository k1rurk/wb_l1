package main

import (
	"fmt"
	"sync"
	"time"
)

//Через передачу значения в канал для выхода
func stopChannel(quit chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-quit:
			return
		default:
			time.Sleep(5 * time.Millisecond)
			fmt.Println("Some text")
		}
	}
}

//Через закрытие канала
func stopChannel2(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for foo := range ch {
		fmt.Println(foo)
	}
}

func main() {

	wg := new(sync.WaitGroup)

	quit := make(chan bool)
	ch := make(chan int)
	wg.Add(1)
	go stopChannel(quit, wg)
	wg.Add(2)
	go stopChannel2(ch, wg)
	time.Sleep(2 * time.Second)
	quit <- true
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	wg.Wait()
	fmt.Println("Stopped")

}
