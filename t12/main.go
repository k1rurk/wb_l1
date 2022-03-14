package main

import "fmt"

func main() {

	str := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]bool)

	for _, v := range str {
		if _, ok := set[v]; !ok {
			set[v] = true
		}
	}

	for i := range set {
		fmt.Println(i)
	}

}
