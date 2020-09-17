package structs

import "math"

// Shape can calculate the area
type Shape interface {
	Area() float64
}

// Rectangle defines a rectange.
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter calculates the perimeter of a rectangle.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area calculates the area of a rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle defines a circle
type Circle struct {
	Radius float64
}

// Area calculates the area of a circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle defines a triangle with base and height.
type Triangle struct {
	Base   float64
	Height float64
}

// Area calculates the area of a triangle
func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}
