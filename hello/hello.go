package main

import (
	"fmt"

	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greeting to humans: ")
	log.SetFlags(0)
	message, err := greetings.Hello("Ayofe")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
