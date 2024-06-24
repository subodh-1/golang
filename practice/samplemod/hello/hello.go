package main

import (
	"fmt"
	"github.com/subodh-1/gopractice/practice/samplemod/greetings"
)

func main(){
	var name string
    fmt.Println("Please enter your name:")
    fmt.Scanf("%s", &name)
    msg := greetings.MorningGreetings(name)
    fmt.Println(msg)
}