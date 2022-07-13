package main

import "math"

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	area() float64
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}
