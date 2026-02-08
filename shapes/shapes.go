package main

import (
	"math"

)


//Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

//Area of the Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

//Perimeter of the Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

//Scale the Rectangle by a factor
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

//Circle
type Circle struct {
	Radius float64
}

//Area of the Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//Perimeter of the Circle
func (c Circle) Perimeter() float64 {
	return math.Pi * 2 * c.Radius
}

//Scale the Rectangle by a Circle
func (c *Circle) Scale(factor float64) {
	c.Radius *= factor
}

//Triangle
type Triangle struct {
	Base   float64
	Height float64
}

//Area of the Triangle
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

//Perimeter of the Triangle
func (t Triangle) Perimeter() float64 {
	return 3 * t.Base
}

//Scale the Triangle by a factor
func (t *Triangle) Scale(factor float64) {
	t.Base *= factor
	t.Height *= factor
}
