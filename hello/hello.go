package main

import (
	"fmt"

	"github.com/Dawreen/go_learning/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Lucia")
	fmt.Println(message)
}
