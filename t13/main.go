package main

import "fmt"

func main() {
	var a, b int64
	fmt.Printf("Введите первое число: ")
	fmt.Scanf("%d\n", &a)
	fmt.Printf("Введите второе число: ")
	fmt.Scanf("%d\n", &b)

	fmt.Printf("a = %v b = %v\n", a, b)
	a = a ^ b
	b = b ^ a
	a = a ^ b
	fmt.Printf("a = %v b = %v\n", a, b)

}
