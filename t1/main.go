package main

import "fmt"

type cm int
type kg float32

type Action struct {
	*Human
}

type Human struct {
	colorEye string
	age      uint
	height   cm
	weight   kg
}

func (a *Action) getOlder() {
	a.age++
}

func (a *Action) sayAboutYourself() {
	fmt.Printf("Цвет глаз = %v\nВес = %v\nРост = %v\nВозраст = %v\n\n\n",
		a.colorEye,
		a.weight,
		a.height,
		a.age,
	)
}

func NewHuman(colorEye string, age uint, height cm, weight kg) *Human {
	return &Human{
		colorEye: colorEye,
		age:      age,
		height:   height,
		weight:   weight,
	}
}

func (a *Action) goGym() {
	a.weight--
}

func main() {
	human := NewHuman("green", 20, 184, 86.6)
	action := Action{human}
	action.sayAboutYourself()
	action.goGym()
	action.getOlder()
	action.sayAboutYourself()
}
