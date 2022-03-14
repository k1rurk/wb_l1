package main

import (
	"fmt"
	"sync"
)

func main() {
	sum := 0
	mx := new(sync.Mutex)
	arr := [5]int{2, 4, 6, 8, 10}
	wg := new(sync.WaitGroup)
	wg.Add(5)
	for i := range arr {
		go func(i int) {
			mx.Lock()
			defer mx.Unlock()
			defer wg.Done()
			sum += i * i
		}(i)
	}
	wg.Wait()
	fmt.Println(sum)
}
