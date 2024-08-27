// interface example
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
	Perimeter() float32
}

type Rectangle struct {
	Height float32
	Width  float32
}
type Circle struct {
	Radius float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Height + r.Width)
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float32 {
	return 2 * (math.Pi * c.Radius)
}

func measure(s Shape) {
	fmt.Printf("Shape Type: %T", s)
	fmt.Println("\nArea: ", s.Area())
	fmt.Println("Perimeter: ", s.Perimeter())
}

func main() {
	r := Rectangle{Height: 10, Width: 15}
	c := Circle{Radius: 10}
	measure(r)
	measure(c)
}
