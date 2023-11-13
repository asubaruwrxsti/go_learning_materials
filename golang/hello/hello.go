package main

import (
	"fmt" // Package implementing formatted I/O.
	"rsc.io/quote"
	"example.com/greetings"
	"log"
)

// the directory structure is important
// after creating the module, run go mod tidy to add the dependencies to go.mod
// run the program with go run . or go run hello.go

func main() { // main function executes by default
	// configure log
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
	fmt.Println(quote.Go())
}
