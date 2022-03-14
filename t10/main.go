package main

import "fmt"

func main() {

	tempertures := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := map[int][]float32{}
	for i := -20; i <= 30; i += 10 {
		var tempArr []float32
		for _, t := range tempertures {
			if float32(i-10) < t && t <= float32(i+10) {
				check := false
				for _, tGroup := range groups[i-10] {
					if tGroup == t {
						check = true
					}
				}
				if !check {
					tempArr = append(tempArr, t)
				}
			}
		}
		if len(tempArr) != 0 {
			groups[i] = tempArr
		}
	}

	fmt.Println(groups)
}
