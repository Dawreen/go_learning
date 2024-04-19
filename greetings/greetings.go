package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name was received, return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos retruns a map that associates each of the named people with a greeting.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// random greeting
func randomFormat() string {
	// All the possible formats
	formats := []string{
		"Hi, %v. Welcome!\nDan is tring to learn GO...\nHe gave up on engliss some time ago",
		"Hi, %v. Welcome!",
		"Nice to see %v",
		"Better luck next time! %v, please LEAVE!",
	}

	// Return a randomly selected message from the formats
	return formats[rand.Intn(len(formats))]
}
