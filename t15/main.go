package main

import (
	"fmt"
	"strconv"
)

var justString string

func createHugeString(a int) string {
	var str string
	for i := 0; i < a; i++ {
		str += strconv.FormatInt(int64(i), 10)
	}
	return str
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
	fmt.Println(justString)
}

func main() {
	someFunc()
}
