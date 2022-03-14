package main

import (
	"fmt"
	"sync"
)

type MyMap struct {
	mx sync.RWMutex
	mp map[int]int
}

func NewMap() *MyMap {
	return &MyMap{
		mp: make(map[int]int),
	}
}

func (m *MyMap) Set(key int, value int) {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.mp[key] = value
}

func (m *MyMap) Get(key int) (int, bool) {
	m.mx.RLock()
	defer m.mx.RUnlock()
	val, ok := m.mp[key]
	return val, ok
}

func main() {

	myMap := NewMap()

	wg := new(sync.WaitGroup)
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			myMap.Set(i, i+1)
		}(i)
	}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			v, _ := myMap.Get(i)
			fmt.Println(v)
		}(i)
	}
	wg.Wait()

}
