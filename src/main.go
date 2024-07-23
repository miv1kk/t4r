package main

import (
	"errors"
	"fmt"

	"rsc.io/quote"
)

func print_and_return(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("%v", name)
	return message, nil
}

func main() {
	fmt.Println(print_and_return("Me"))

	fmt.Println("Hello, world!")
	fmt.Println("Hello, world!")

	fmt.Println(quote.Go())
}
