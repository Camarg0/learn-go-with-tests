package structsmethodsinterfaces

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) CalculateArea() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) CalculateArea() float64 {
	return c.Radius * c.Radius * math.Pi
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) CalculateArea() float64 {
	return t.Base * t.Height / 2
}

type Shape interface {
	CalculateArea() float64
}

func CalculatePerimeter(width float64, height float64) float64 {
	return 2 * (width + height)
}
