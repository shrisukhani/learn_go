package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float64
}

type square struct {
	side float64
}

type triangle struct {
	sideOne   float64
	sideTwo   float64
	sideThree float64
}

func (t triangle) getArea() float64 {
	semiperimeter := (t.sideOne + t.sideTwo + t.sideThree) / 2
	return math.Sqrt(semiperimeter * (semiperimeter - t.sideOne) * (semiperimeter - t.sideTwo) * (semiperimeter - t.sideThree))
}

func (s square) getArea() float64 {
	return s.side * s.side
}

func printArea(s shape) {
	fmt.Println("Area of shape:", s.getArea())
}

func main() {
	printArea(triangle{4, 4, 4})
	printArea(square{4})
}
