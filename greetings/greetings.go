package greetings

import (
	"fmt"
)

func Hello(name string) string {
	var message string
	message = fmt.Sprintf("Hello %v, you are welcome in the module", name)
	return message
}