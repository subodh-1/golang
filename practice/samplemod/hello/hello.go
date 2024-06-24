package main

import (
	"fmt"
	"github.com/subodh-1/gopractice/practice/samplemod/greetings"
)

func main(){
	name := "Subodh Choure"
	msg := greetings.MorningGreetings(name)
	fmt.Print(msg)
}