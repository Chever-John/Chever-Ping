package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Triangle struct {
	Width  float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Width * t.Height / 2
}

func PrintArea(s Shape) {
	fmt.Println(s.Area())
}

func main() {
	r := Rectangle{Width: 3, Height: 4}
	c := Circle{Radius: 5}
	t := Triangle{Width: 3, Height: 4}

	PrintArea(r)
	PrintArea(c)
	PrintArea(t)
}
