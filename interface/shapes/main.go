package main

import (
	"fmt"
)

// create a shape square - sideLength (float64)
// create a shape triangle - height (float64), base (float64)
// function that gets the area of a triangle - .5 * base * height
// function that get the area of the square - sideLength * sideLength
// interface for shape
// shape printArea function

type square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	height float64
}

type shape interface {
	getArea() float64
}

func main() {
	s := square{sideLength: 10}
	t := triangle{base: 10, height: 5}

	printArea(s)
	printArea(t)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}
