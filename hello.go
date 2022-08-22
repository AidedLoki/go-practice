package main

import (
	"fmt"
	"log"

	"rsc.io/quote"

	"net/http"

	"github.com/AidedLoki/greetings"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	names := []string{"Pikey", "Samantha", "Chase"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Got a request !!")

		log.SetPrefix("Greetings: ")
		log.SetFlags(0)

		message, err := greetings.Hellos(names)

		if err != nil {

			fmt.Fprintf(w, "Hello, you've requested error: %s\n", err)
			//log.Fatal(err)
		}

		fmt.Println(message)
		fmt.Fprintf(w, "Hello, you've requested: %s\n", message)
	})
	http.ListenAndServe(":80", nil)

}
