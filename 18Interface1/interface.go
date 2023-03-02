package main

import (
	"fmt"
	"math"
)

func main() {
	shapes := []Shape{
		Circle{1.0},
		Square{1.772453},
		Rectangle{5, 10},
		Triangle{10, 4, 7},
	}
	for _, v := range shapes {
		fmt.Println(v, "\tArea:", v.Area())
		if t, ok := v.(Triangle); ok {
			fmt.Println("Angles:", t.Angles())
		}
	}
}

type Shape interface {
	Area() float64
}
type Circle struct {
	Radius float64
}
type Triangle struct {
	A, B, C float64 // lengths of the sides of a triangle.
}
type Rectangle struct {
	A, B float64
}
type Square struct {
	A float64
}

func (t Circle) Area() float64 {
	return math.Pi * t.Radius * t.Radius
}

// Heron's Formula for the area of a triangle
func (t Triangle) Area() float64 {
	p := (t.A + t.B + t.C) / 2.0 // perimeter half
	return math.Sqrt(p * (p - t.A) * (p - t.B) * (p - t.C))
}
func (t Rectangle) Area() float64 {
	return t.A * t.B
}

func (t Square) Area() float64 {
	return t.A * t.A
}

func (t Circle) String() string {
	return fmt.Sprint("Circle (Radius: ", t.Radius, ")")
}
func (t Triangle) String() string {
	return fmt.Sprint("Triangle (Sides: ", t.A, ", ", t.B, ", ", t.C, ")")
}
func (t Rectangle) String() string {
	return fmt.Sprint("Rectangle (Sides: ", t.A, ", ", t.B, ")")
}
func (t Square) String() string {
	return fmt.Sprint("Square (Sides: ", t.A, ")")
}

func (t Triangle) Angles() []float64 {
	return []float64{angle(t.B, t.C, t.A), angle(t.A, t.C, t.B), angle(t.A, t.B, t.C)}
}
func angle(a, b, c float64) float64 {
	return math.Acos((a*a+b*b-c*c)/(2*a*b)) * 180.0 / math.Pi
}
