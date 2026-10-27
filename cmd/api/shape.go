package main

import "math"

type Shape interface {
	Perimeter() float64
}

type Rectangle struct{
	width float64
	height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

type Circle struct{
	radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * c.radius * math.Pi
}

