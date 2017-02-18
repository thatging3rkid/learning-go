package main

import (
	"fmt"
)

func main() {
	var big uint64
	prime := true
	line_prints := 0

	for big = 1; big != 0; big += 1 {
		var i uint64
		for i = 2; i < big; i += 1 {
			if (big % i) == 0 {
				prime = false
				break
			}
		}
		if prime {
			if line_prints > 12 {
				fmt.Println()
				line_prints = 0
			}
			fmt.Print("Prime: " + fmt.Sprint(big) + " ")
			line_prints += 1
		}
		prime = true
	}
}