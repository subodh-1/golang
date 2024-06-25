package greetings

import (
	"fmt"
	"errors"
)

func checkNameEmpty(name string) error {	
	if name == "" {
		return errors.New("Name can not be empty!!")
	}
	return nil
}

func Greet(name string, time string) (string, error) {
	var msg string
	err := checkNameEmpty(name)
	if err != nil {
		return "", err
	}
	switch time {
		case "morning": 
			msg = MorningGreetings(name)			
		case "afternoon":
			msg = AfternoonGreetings(name)			
		case "evening":
			msg = EveningGreetings(name)			
		case "night":
			msg = NightGreetings(name)			
		default:
			return "Sorry, no time matches", err
	}
	return msg, err
}

func MorningGreetings(name string) string {
	return fmt.Sprintf("Hello %s !!, Good morning!!", name)
}
func AfternoonGreetings(name string) string {
	return fmt.Sprintf("Hello %s !!, Good afternoon!!", name)
}
func EveningGreetings(name string) string {
	return fmt.Sprintf("Hello %s !!, Good evening!!", name)
}
func NightGreetings(name string) string {
	return fmt.Sprintf("Hello %s !!, Good night!!", name)
}