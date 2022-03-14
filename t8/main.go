package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int64
	var digit int
	var isZero int
	fmt.Printf("Введите число: ")
	_, err := fmt.Scanf("%d\n", &number)
	if err != nil {
		return
	}
	fmt.Printf("Введите разряд: ")
	_, err = fmt.Scanf("%d\n", &digit)
	if err != nil {
		return
	}
	fmt.Printf("Ноль или единица: ")
	_, err = fmt.Scanf("%d\n", &isZero)
	if err != nil {
		return
	}
	if isZero != 0 && isZero != 1 {
		fmt.Println("Вы ввели ни ноль, ни единицу")
		return
	}

	makeBit(number, digit, isZero != 1)

}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func makeBit(number int64, digit int, isZero bool) {

	if digit > 64 && digit < 0 {
		fmt.Println("Разряд бита задан неверно")
		return
	}
	fmt.Println("Вы ввели число ", number)
	str := strconv.FormatInt(number, 2)
	fmt.Println("Битовое представление до ", str)
	str = Reverse(str)
	if isZero {
		str = str[:digit] + "0" + str[digit+1:]
	} else {
		str = str[:digit] + "1" + str[digit+1:]
	}
	str = Reverse(str)
	fmt.Println("Битовое представление после ", str)
	i, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Числовое представление ", i)
}
