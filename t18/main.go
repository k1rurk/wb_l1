package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type counter struct {
	mx sync.Mutex
	c  int
}

func (c *counter) increment() {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.c++
}

func getRandNum(min, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	c := new(counter)
	rand.Seed(time.Now().UnixNano())
	n := getRandNum(10, 100)
	wg := new(sync.WaitGroup)

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			c.increment()
		}()
	}
	wg.Wait()
	fmt.Println("Счетчик = ", c.c)

}
