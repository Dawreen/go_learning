package main

import (
	"fmt"

	"go_learning/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Lucia")
	fmt.Println(message)
}
