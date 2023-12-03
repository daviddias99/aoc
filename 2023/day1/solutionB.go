package day1

import (
	"aoc/2023/utils"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func valueFromString(str string) (int, error) {

	if strings.HasPrefix(str, "zero") {
		return 0, nil
	} else if strings.HasPrefix(str, "one") {
		return 1, nil
	} else if strings.HasPrefix(str, "two") {
		return 2, nil
	} else if strings.HasPrefix(str, "three") {
		return 3, nil
	} else if strings.HasPrefix(str, "four") {
		return 4, nil
	} else if strings.HasPrefix(str, "five") {
		return 5, nil
	} else if strings.HasPrefix(str, "six") {
		return 6, nil
	} else if strings.HasPrefix(str, "seven") {
		return 7, nil
	} else if strings.HasPrefix(str, "eight") {
		return 8, nil
	} else if strings.HasPrefix(str, "nine") {
		return 9, nil
	}

	return 0, errors.New("NAN")
}

func SolveB() {
	fmt.Println("Day 1 Problem B")

	lines := utils.ReadLines("day1/input/b.input")

	result := 0
	values := make([]int, len(lines))

	for lineno, line := range lines {
		found := false
		digit := 0
		lastdigit := 0
		var err error

		for i, char := range line {
			digit, err = strconv.Atoi(string(char))
			if err == nil {
				lastdigit = digit
				if !found {
					found = true
					values[lineno] += digit * 10
				}

				continue
			}

			digit, err = valueFromString(line[i:])

			if err == nil {
				lastdigit = digit
				if !found {
					found = true
					values[lineno] += digit * 10
				}

				continue
			}
		}

		values[lineno] += lastdigit
		result += values[lineno]
	}

	// fmt.Println(values)
	fmt.Println(result)
}
