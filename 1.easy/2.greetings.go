package main

import "fmt"

func main() {
	fmt.Print("Please enter your name: ")
	input := ""
	fmt.Scanln(&input)
	fmt.Println("Hello, " + input)
}