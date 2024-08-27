package main

import (
	"fmt"
)

/**
Go example to calcualate area and perimeter with the use of interface and struct
*/

type Shape interface {
	Area() float32
	Perimeter() float32
}

type Rectangle struct {
	Height float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return (r.Height * r.Width)
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Height + r.Width)
}
func main() {
	var s Shape = Rectangle{Height: 5, Width: 15}
	fmt.Println("Area of rectangle: ", s.Area())
	fmt.Println("Perimeter of rectange: ", s.Perimeter())
}
