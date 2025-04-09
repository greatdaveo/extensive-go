package mylib

import (
	"fmt"
	"math"
)

// An interface is a way of defining a set of method that other types must implement
// in order for them to be considered the type that interface is of.

type geometry interface {
	area() float64
	perim() float64
}

// rect
type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perim() float64 {
	return 2 * (r.height + r.width)
}

// rect 1

type rect1 struct {
	width, height float64
}

func (r rect1) area() float64 {
	return r.height * r.width
}

func (r rect1) perim() float64 {
	return 2 * (r.height + r.width)
}

// circle
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) diameter() float64 {
	return 2 * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func InterfaceProcess() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	r1 := rect1{width: 3, height: 4}

	measure(r)
	measure(c)
	measure(r1)

	myPrinter(1, "Dave", 35.8, true)

	printType(9)
	printType("John Doe")
	printType(false)

}

func printType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Type: Int")
	case string:
		fmt.Println("Type: String")
	default:
		fmt.Println("Type: Unknown")
	}

}

func myPrinter(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v)
	}
}
