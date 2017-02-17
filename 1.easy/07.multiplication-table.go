package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter the max number for the multiplication table: ")
	var max int
	fmt.Scanf("%d", &max)



	for i := 0; i < max + 1; i += 1 {
		for k := 0; k < max + 1; k += 1 {
			fmt.Printf("%03d | ", i * k)
		}
		fmt.Print("\n")
	}

}
