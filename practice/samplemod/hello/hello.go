package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/subodh-1/gopractice/practice/samplemod/greetings"
)

func main() {
	log.SetPrefix("Greetings error:")
	log.SetFlags(0)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println("Please enter the time of day (morning, afternoon, evening, night):")
	time, _ := reader.ReadString('\n')
	time = strings.TrimSpace(time)
	greet, err := greetings.Greet(name, time)
	if err != nil {
		//fmt.Println("Error while performing the operation")
		log.Fatal(err)
	} else {
		fmt.Println(greet)
	}

}
