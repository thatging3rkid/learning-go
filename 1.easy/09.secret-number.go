package main

import (
	"fmt"
	"syscall"
	"time"
	"math/rand"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

func main() {

	var min int
	var max int

	fmt.Print("Enter a minimum number: ")
	fmt.Scanf("%d", &min)
	fmt.Print("Enter a maximum number: ")
	fmt.Scanf("%d", &max)

	if min > max {
		fmt.Println("min is greater than max\nExiting...")
		syscall.Exit(1)
	}

	target := random(min, max)
	input := 0

	for true {
		fmt.Print("Guess a number: ")
		fmt.Scanf("%d", &input)

		if target == input {
			fmt.Print("You got it!")
			break
		} else {
			if input < target {
				fmt.Print("Too low...\n")
			}
			if input > target {
				fmt.Print("Too high...\n")
			}
		}


	}


}
