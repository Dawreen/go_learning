package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!\nDan is tring to learn GO...\nHe gave up on engliss some time ago", name)
	return message
}
