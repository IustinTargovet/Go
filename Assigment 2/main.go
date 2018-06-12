package main

import (
	"fmt"
)

type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	var t triangle
	t.base = 10
	t.height = 10
	printArea(t)
	var s square
	s.sideLength = 5
	printArea(s)
}
