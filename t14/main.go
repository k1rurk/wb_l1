package main

import "fmt"

func main() {
	ch := make(chan bool)
	interf := []interface{}{"a", 12, true, ch}

	for _, val := range interf {
		switch v := val.(type) {
		case int:
			fmt.Printf("Integer: %v\n", v)
		case bool:
			fmt.Printf("Float64: %v\n", v)
		case string:
			fmt.Printf("String: %v\n", v)
		case chan bool:
			fmt.Printf("Channel bool: %v\n", v)
		default:
			fmt.Printf("None type")
		}
	}
}
