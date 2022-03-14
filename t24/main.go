package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

func NewPoint(x, y int) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func distancePoint(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow(float64(p2.x-p1.x), 2) + math.Pow(float64(p2.y-p1.y), 2))
}

func main() {
	p1 := NewPoint(-1, 3)
	p2 := NewPoint(6, 2)

	dist := distancePoint(p1, p2)
	fmt.Println("Расстояние = ", dist)

}
