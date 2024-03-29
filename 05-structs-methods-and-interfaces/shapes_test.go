package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, wanted %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{width: 12, height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{base: 12, height: 6}, hasArea: 36.0},
	}

	// using tt.name from the case to use it as the `t.Run` test name
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shape.area() != tt.hasArea {
				t.Errorf("%#v got %.2f, wanted %.2f", tt.name, tt.shape.area(), tt.hasArea)
			}
		})
	}
}
