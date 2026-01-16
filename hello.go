package main

import (
	"fmt"
	"log"

	"github.com/Sahil2004/greetings"
	"rsc.io/quote"
)

func main() {
	log.SetPrefix("log: ")
	log.SetFlags(0)
	fmt.Println("Hello, World!");
	fmt.Println(quote.Go());
	message, err := greetings.Hello("Sahil")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{"Alice", "Bob", "Charlie"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}