package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print("Enter a number: ")
	var input int
	fmt.Scanf("%d", &input)
	sum := 0
	for i := 0; i <= input; i += 1 {
		if ((i % 3) == 0) || ((i % 5) == 0) {
			sum += i
		}
	}
	fmt.Println("Sum of 1 to " + strconv.Itoa(input) + " with only numbers that are divisible by 3: " + strconv.Itoa(sum))
}
