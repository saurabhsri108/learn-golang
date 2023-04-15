package main

import (
	"fmt"
)

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func assignment1() {
	sq := square{sideLength: 10}
	t := triangle{height: 4, base: 5}
	printArea(sq, "square")
	printArea(t, "triangle")
}

func printArea(s shape, figure string) {
	fmt.Printf("Area of %s with dimensions %+v: %v\n", figure, s, s.getArea())
}
