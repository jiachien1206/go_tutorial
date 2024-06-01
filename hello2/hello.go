// caller of greetings module

package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
    message := greetings.Hello("Nini")
		fmt.Println(message)
}