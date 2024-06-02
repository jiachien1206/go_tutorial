// caller of greetings module

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

	// correct input
	message, err := greetings.Hello("Nini")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	// with error
	m, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)
}