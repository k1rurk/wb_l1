package main

import "fmt"

type Set struct {
	s map[int]bool
}

func NewSet() *Set {
	return &Set{
		s: map[int]bool{},
	}
}

func (set *Set) fillSet(arr []int) {
	for _, v := range arr {
		set.s[v] = true
	}
}

func (set *Set) intercept(set2 *Set) *Set {

	setTemp := NewSet()

	for v := range set.s {
		if _, ok := set2.s[v]; ok {
			setTemp.s[v] = true
		}
	}

	return setTemp
}

func main() {
	set1 := NewSet()
	set2 := NewSet()

	arr1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 := []int{2, 4, 6, 8, 12, 22}

	set1.fillSet(arr1)
	set2.fillSet(arr2)

	setIntercept := set1.intercept(set2)

	for v := range setIntercept.s {
		fmt.Println(v)
	}
}
