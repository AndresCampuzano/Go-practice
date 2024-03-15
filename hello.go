package main

import (
	"fmt"
	"math"
)

//type ShapeAreas interface {
//	RectangleArea() float64
//	CircleArea() float64
//}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

//func (r Rectangle) RectangleArea() float64 {
//	return r.Width * r.Height
//}

func (r Rectangle) AnotherMethod() string {
	return "Another method\n"
}

type Circle struct {
	Radius float64
}

func (r Circle) Area() float64 {
	return math.Pi * r.Radius * r.Radius
}

//func (r Circle) CircleArea() float64 {
//	return math.Pi * r.Radius * r.Radius
//}

func main() {
	r := Rectangle{
		Width:  10,
		Height: 10,
	}
	fmt.Printf("Area of rectangle %#v is = %g \n", r, r.Area())

	c := Circle{Radius: 50}
	fmt.Printf("Area of circle %#v is = %g \n", c, c.Area())
}
