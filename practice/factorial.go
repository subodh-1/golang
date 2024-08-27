package main

import "fmt"

func main() {
	fmt.Println("Factorial of x")
	fmt.Println(factorial_rec(4))
}

func factorial_rec(x float64) (y float64) {
	if x > 0 {
		y = x * factorial_rec(x-1)
	} else {
		y = 1
	}
	return
}
