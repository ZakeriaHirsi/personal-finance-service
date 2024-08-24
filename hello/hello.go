package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Println(reverse.String("Zak"))
}

func HelloGreeting() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Zak", "John", "Sarah", "Barthalomeow"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
