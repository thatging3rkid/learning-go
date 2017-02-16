package main

import (
	"fmt"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	// Gets user input
	fmt.Print("Enter a number: ")
	var input int
	fmt.Scanf("%d", &input)
	fmt.Print("Would you like the sum or multiply (s or m): ")
	var add_mode string
	fmt.Scanf("%s", &add_mode)

	// Determines if it should be in sum or multiply mode
	sum_mode := false
	if strings.ToLower(add_mode) == "s" {
		sum_mode = true
	} else if strings.ToLower(add_mode) == "m" {
		sum_mode = false
	} else {
		fmt.Println("Unrecognized input. Exiting...")
		syscall.Exit(1)
	}

	// Does the math
	sum := 0
	for i := 0; i <= input; i += 1 {
		if sum_mode {
			sum += i
		} else {
			if i == 0 {
				sum = 1 // need to set to one because 0 * x = 0
			} else {
				sum *= i
			}
		}
	}

	// Prints the result
	var output_word string
	if sum_mode {
		output_word = "Sum"
	} else {
		output_word = "Multiple"
	}
	fmt.Println(output_word + " of 1 to " + strconv.Itoa(input) + ": " + strconv.Itoa(sum))
}
