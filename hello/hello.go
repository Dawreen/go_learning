package main

import (
	"fmt"
	"log"

	"go_learning/greetings"
)

func main() {
	// Set properties of the predefined logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings:")
	log.SetFlags(0)

	// A slice of names
	names := []string{"Dan", "Dorin", "Gandalf", "Luke", "Harry"}

	// Request a greeting message.
	messages, err := greetings.Hellos(names)
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console
	fmt.Println(messages)
}
