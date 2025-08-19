package structs

import "math"

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	base   float64
	height float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (t Triangle) Area() float64 {
	return t.base * t.height * 0.5
}
