package main

import "fmt"

type BasicCar interface {
	drive()
	xenonOn()
	xenonOff()
}

type AudiA3 struct {
}

func (a *AudiA3) drive() {
	fmt.Println("AudiA3 rides!")
}

func (a *AudiA3) xenonOn() {
	fmt.Println("AudiA3 xenon on")
}

func (a *AudiA3) xenonOff() {
	fmt.Println("AudiA3 xenon off")
}

type CentralProcessor struct {
}

func (p *CentralProcessor) work(basicCar BasicCar) {
	basicCar.drive()
	basicCar.xenonOn()
	basicCar.xenonOff()
}

type AudiA3Rus struct {
}

func (r *AudiA3Rus) drive() {
	fmt.Println("AudiA3Rus rides!")
}

func (r *AudiA3Rus) basicLightOn() {
	fmt.Println("AudiA3Rus basic light on")
}

func (r *AudiA3Rus) basicLightOff() {
	fmt.Println("AudiA3Rus basic light off")
}

type CarAdapter struct {
	audiA3Rus *AudiA3Rus
}

func (a *CarAdapter) drive() {
	a.audiA3Rus.drive()
}

func (a *CarAdapter) xenonOn() {
	a.audiA3Rus.basicLightOn()
}

func (a *CarAdapter) xenonOff() {
	a.audiA3Rus.basicLightOff()
}

func main() {
	audiA3 := new(AudiA3)
	cp := CentralProcessor{}
	cp.work(audiA3)

	audiA3Rus := new(AudiA3Rus)
	adapterRus := &CarAdapter{audiA3Rus}
	cpRus := CentralProcessor{}
	cpRus.work(adapterRus)
}
