package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input int
	fmt.Print("Enter a year: ")
	fmt.Scanf("%d", &input)

	printed := 0
	for true {
		if (input % 400 == 0) || ((input % 4)  == 0 && (input % 100 != 0)) {
			printed += 1
			fmt.Println("Leap year on: " + strconv.Itoa(input))
			if printed == 20 {
				break
			}
		}

		input += 1
	}
}