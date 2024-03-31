package main

import "fmt"

type gasEngine struct {
	mgp       int
	gallons   int
	ownerInfo owner
}

type owner struct {
	name string
}

type ElectricEngine struct {
	mkw        int
	chargeleft int
}

func (e gasEngine) mileleft() int {
	return e.gallons * e.mgp
}

func (e ElectricEngine) mileleft() int {
	return e.chargeleft * e.mkw
}

type engine interface {
	mileleft() int
}

func canmakeIt(e engine, miles int) {
	if miles < e.mileleft() {
		fmt.Println("Will make it")
	} else {
		fmt.Println("need Fuel")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{12, 2, owner{"Alex"}}
	fmt.Println(myEngine.mgp, myEngine.gallons, myEngine.ownerInfo.name)
	fmt.Println(myEngine.mileleft())

	canmakeIt(myEngine, 100)
}
