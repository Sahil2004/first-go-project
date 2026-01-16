package main

import (
	"fmt"
	"log"
	"rsc.io/quote"
	"github.com/Sahil2004/greetings"
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
}