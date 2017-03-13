package main

import (
	"fmt"
	"strconv"
	//"strings"
)

func parseIntegerList(input string) []int {
	var output []int
	last_comma := 0
	output_head := 0

	for i := 0; i < len(input); i += 1 {
		if input[i] == ',' {
			fmt.Print(input[last_comma:i])
			fmt.Print("`")
			output_head, _ = strconv.Atoi(input[last_comma:i])
			output = append(output, output_head)
			last_comma = i + 1
			output_head += 1
		} else {
			continue
		}

	}

	return output
}

func main() {
	fmt.Print("Enter an integer list:")
	var input string
	fmt.Scanf("%s", &input)
	input_array := parseIntegerList(input)
	fmt.Println(input_array)
}