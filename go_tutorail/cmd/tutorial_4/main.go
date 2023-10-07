package main

import "fmt"

// Struc and interface

type gasEngine struct {
	kmpl   uint8 // km / L
	litres uint8
}

func (e gasEngine) kmLeft() uint8 {
	return e.litres * e.kmpl
}

type electricEngine struct {
	kmkwh uint8 //
	kwh   uint8
}

func (e electricEngine) kmLeft() uint8 {
	return e.kwh * e.kmkwh
}

type engine interface {
	kmLeft() uint8
}

func canMakeIt(e engine, km uint8) {
	if km <= e.kmLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("You will need to fuel up!")
	}

}

func main() {
	var myEngine = gasEngine{kmpl: 10, litres: 22}
	var myElecEngine = electricEngine{kmkwh: 10, kwh: 22}
	myEngine.kmpl = 150
	fmt.Println(myEngine.kmpl, myEngine.litres)
	fmt.Printf("Total of km left is tank: %v\n", myEngine.kmLeft())

	canMakeIt(myEngine, 200)
	canMakeIt(myElecEngine, 200)
}
