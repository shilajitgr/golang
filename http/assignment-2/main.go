package main

import (
	"fmt"
	"math"
)

type triangle struct {
	height float64
	base   float64
}

type square struct {
	side float64
}

type shape interface {
	getArea() float64
}

func main() {
	t := triangle{2, 4.8}
	sq := square{3}

	printArea(t)
	printArea(sq)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2
}

func (s square) getArea() float64 {
	return math.Pow(s.side, 2)
}
