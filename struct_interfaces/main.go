package main

import "fmt"

type engine interface {
	milesLeft() uint8
}

type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner
	int
}

type electiceEngine struct {
	mpkwh uint8
	kwh uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons*e.mpg
}
func (e electiceEngine) milesLeft() uint8 {
	return e.kwh*e.mpkwh
}

func canMakeIt(e engine, miles uint8) bool {
	if e.milesLeft() == miles {
		return true
	} else {
		return false
	}
}

type owner struct {
	name string
}

func main() {
	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15, owner: owner{name: "Estif"}, int: 2}
	myEngine.gallons += 4
	fmt.Println(myEngine.name, myEngine.int)
	fmt.Println(myEngine.gallons, myEngine.mpg)

	// annonymus struct
	var car = struct{
		name string
		model string
		wheels int
	}{"max", "Toyota", 4}
	fmt.Println(car)
}