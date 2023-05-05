package greetings

import (
	"fmt"

	"github.com/billylkc/greetings/src/utils"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	sth := utils.ReturnSth()
	message := fmt.Sprintf("Hi, %v. Welcome! %s", name, sth)
	return message
}
