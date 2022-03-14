package main

import (
	"fmt"
	"wb_l1/t16/quickSort"
)

func binarySearch(arr []int, key int) (int, bool) {
	if len(arr) == 0 {
		return -1, false
	}
	fmt.Println("Сортируем его...")
	quickSort.QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
	l := 0
	r := len(arr) - 1
	search := -1
	for l <= r {
		mid := l + (r-l)/2
		if key == arr[mid] {
			search = mid
			return search, true
		}
		if arr[mid] < key {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return search, false
}

func main() {
	arr := []int{10, 20, -22, 1, 0, -2, 4, -100, -3, 40, 32, -9, 10, 34, 11}
	key := -100
	fmt.Printf("Дан массив %v\nИщем значение %v\n", arr, key)
	if v, ok := binarySearch(arr, key); ok {
		fmt.Printf("Значение найдено, индекс = %d\n", v)
	} else {
		fmt.Println("Значение не найдено")
	}
}
