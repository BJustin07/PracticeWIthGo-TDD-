package structMethodsAndArrays

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	base, height float64
}

func (t Triangle) Area() float64 {
	return (t.base * t.height) * 0.5
}
