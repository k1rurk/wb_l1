package main

import "fmt"

func deleteByIndexSlice(arr []string, i int) []string {
	// Выполнить сдвиг arr[i+1:] влево на один индекс
	copy(arr[i:], arr[i+1:])

	// Удаляем последний элемент
	arr = arr[:len(arr)-1]

	return arr
}

func main() {
	arr := []string{"A", "B", "C", "D", "E"}

	arr = deleteByIndexSlice(arr, 2)

	fmt.Println(arr)
}
