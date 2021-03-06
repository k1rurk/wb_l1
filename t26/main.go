package main

import (
	"fmt"
	"strings"
)

func checkUniqueChars(str string) bool {
	str = strings.ToLower(str)
	tempMap := make(map[rune]interface{})
	for _, v := range str {
		if _, ok := tempMap[v]; ok {
			return false
		} else {
			tempMap[v] = new(interface{})
		}
	}
	return true
}

func main() {

	s1, s2, s3 := "abcd", "abCdefAaf", "abCdefAaf"
	s4, s5 := "hello", "helo"

	fmt.Printf("Строка %v = %v\n", s1, checkUniqueChars(s1))
	fmt.Printf("Строка %v = %v\n", s2, checkUniqueChars(s2))
	fmt.Printf("Строка %v = %v\n", s3, checkUniqueChars(s3))
	fmt.Printf("Строка %v = %v\n", s4, checkUniqueChars(s4))
	fmt.Printf("Строка %v = %v\n", s5, checkUniqueChars(s5))

}
