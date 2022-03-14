package main

import (
	"fmt"
	"wb_l1/t16/quickSort"
)

func main() {
	arr := []int{10, 20, -22, 1, 0, -2, 4, -100, -3, 40, 32, -9, 10, 34, 11}
	quickSort.QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
