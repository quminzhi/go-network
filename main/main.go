package main

import (
	"fmt"
	"log"
	"matthewq/tcp"
)

func main() {
	log.SetPrefix("tcp:")
	log.SetFlags(0)

	names := []string{"Matthew", "Ally", "Brown"}
	// Get a greeting message and print it.
	messages, err := tcp.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	for _, message := range messages {
		fmt.Println(message)
	}
}
