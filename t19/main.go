package main

import (
	"fmt"
)

func reverseCharsInStr(str *string) {
	buf := []rune(*str)

	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}

	*str = string(buf)
}

func main() {
	str := "snow dog sun"
	reverseCharsInStr(&str)
	fmt.Println(str)

	str2 := "главрыба"
	reverseCharsInStr(&str2)
	fmt.Println(str2)
}
