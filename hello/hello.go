package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request greeting
	message, err := greetings.Hello("Gladys")
	// If error was returned, print to console and exit program
	if err != nil {
		log.Fatal(err)
	}

	// If no error, print returned msg to console
	fmt.Println(message)
}
