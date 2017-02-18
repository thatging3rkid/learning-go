package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print("Enter the max number for the multiplication table: ")
	var max int
	fmt.Scanf("%d", &max)

	radix := 0
	var temp = max * max * 10
	for temp > 1 {
		temp /= 10
		radix += 1
	}

	for l := 0; l < radix+1; l += 1 {
		fmt.Print(" ")
	}
	fmt.Print("| ")
	for i := 0; i < max+1; i += 1 {
		if i == 0 {
			for j := 0; j < max; j += 1 {
				fmt.Printf("%"+strconv.Itoa(radix)+"d | ", j)
			}
			fmt.Printf("%"+strconv.Itoa(radix)+"d ", max)
			fmt.Print("\n")
		}
		for k := 0; k < max+1; k += 1 {
			if k == 0 {
				fmt.Printf("%"+strconv.Itoa(radix)+"d | ", i)
			}

			if k != max {
				fmt.Printf("%"+strconv.Itoa(radix)+"d | ", i*k)
			} else {
				fmt.Printf("%"+strconv.Itoa(radix)+"d  ", i*k)
			}

		}
		fmt.Print("\n")
	}

}
