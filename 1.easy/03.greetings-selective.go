package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Please enter your name: ")
	input := ""
	fmt.Scanln(&input)
	if strings.ToLower(input) == "alice" || strings.ToLower(input) == "bob" {
		fmt.Println("Hello, " + input)
	} else {
		fmt.Println("Goodbye")
	}

}