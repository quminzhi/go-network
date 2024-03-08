package main

import (
	"fmt"
	"matthewq/tcp"
)

func main() {
	// Get a greeting message and print it.
	message := tcp.Hello("Gladys")
	fmt.Println(message)
}
