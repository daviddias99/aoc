package day1

import (
	"aoc/2023/utils"
	"fmt"
	"strconv"
)

func SolveA() {
	fmt.Println("Day 1 Problem A")

	lines := utils.ReadLines("day1/input/a.input")
	result := 0

	values := make([]int, len(lines))
	for lineno, line := range lines {
		found := false
		digit := 0
		lastdigit := 0
		var err error

		for _, char := range line {
			digit, err = strconv.Atoi(string(char))

			if err == nil {
				lastdigit = digit
				if !found {
					found = true
					values[lineno] += digit * 10
				}
			}
		}

		values[lineno] += lastdigit
		result += values[lineno]
	}

	// fmt.Println(values)
	fmt.Println(result)
}
