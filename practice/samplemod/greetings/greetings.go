package greetings

import (
	"fmt"
)

func MorningGreetings(name string) string {
	return fmt.Spritf("Hello %v !!, Good Morning!!")
}