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
		sum += i
	}
	fmt.Println("Sum of 1 to " + strconv.Itoa(input) + ": " + strconv.Itoa(sum))
}
