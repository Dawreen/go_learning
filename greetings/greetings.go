package greetings

import (
	"errors"

	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name was received, return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!\nDan is tring to learn GO...\nHe gave up on engliss some time ago", name)
	return message, nil
}
