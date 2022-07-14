package main

import "math"

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	base   float64
	height float64
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
	return math.Pi * c.radius * c.radius
}

func (t Triangle) area() float64 {
	return 0.5 * t.base * t.height
}
