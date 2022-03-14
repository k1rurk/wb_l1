package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

func getRandInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func worker(ch chan int, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Println("Worker ", id, " get value ", v)
		time.Sleep(time.Second + time.Duration(getRandInt(1, 3)))
	}
}

func main() {

	var countWorker int
	fmt.Printf("Введите количество воркеров: ")
	fmt.Scanf("%d\n", &countWorker)
	ch := make(chan int, 20)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	wg := new(sync.WaitGroup)

	for i := 0; i < countWorker; i++ {
		wg.Add(1)
		go worker(ch, i, wg)
	}
	rand.Seed(time.Now().UnixNano())
	go func() {
		for {
			i := getRandInt(10, 100)
			select {
			case ch <- i:
				fmt.Println("Sending ", i)
			case <-c:
				fmt.Println("Chanel closed...")
				close(ch)
				return
			}
		}
	}()

	wg.Wait()

}
