package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(40e16)
	b := big.NewInt(5e17)
	r := new(big.Int)

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	r.Add(a, b)
	fmt.Println("a + b = ", r)

	r.Div(b, a)
	fmt.Println("b / a = ", r)

	r.Sub(a, b)
	fmt.Println("a - b = ", r)

	r.Mul(a, b)
	fmt.Println("a * b = ", r)

}
