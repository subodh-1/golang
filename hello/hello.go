package main

import (
	"fmt"
	"example.com/greetings"
)

func main() {
	var name string
	fmt.Println("Enter your name")
	fmt.Scanf("%s", &name)
	msg := greetings.Hello(name)
	fmt.Println(msg)
}