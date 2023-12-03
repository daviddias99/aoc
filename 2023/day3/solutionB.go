package day3

import (
	"aoc/2023/utils"
	"fmt"
	"strconv"
	"strings"
)

var numbers = "0123456789"

func parseNumbetAt(board *[]string, y int, x int) int {
	i := x
	var k int
	for k = x; k >= 0; k-- {
		char := (*board)[y][k]

		if !(char >= '0' && char <= '9') {
			i = k + 1
			break
		} else {
			i = k
		}
	}

	number_str := ""

	for j := i; j < len((*board)[y]); j++ {
		char := (*board)[y][j]
		if char >= '0' && char <= '9' {
			number_str += string(char)
		} else {
			break
		}
	}

	result, err := strconv.Atoi(number_str)

	if err != nil {
		panic("Could not convert number")
	}

	return result
}

func isAdjacentToNumber(board *[]string, y int, x int) (bool, int) {

	// | a b c |
	// | d X e |
	// | f g h |

	var adjacent_numbers []int

	if y != 0 {

		left := false
		center := false
		right := false

		// a
		if x != 0 && strings.Contains(numbers, string((*board)[y-1][x-1])) {
			left = true
		}

		// b
		if strings.Contains(numbers, string((*board)[y-1][x])) {
			center = true
		}

		// c
		if x != (len((*board)[0])-1) && strings.Contains(numbers, string((*board)[y-1][x+1])) {
			right = true
		}

		if center {
			adjacent_numbers = append(adjacent_numbers, parseNumbetAt(board, y-1, x))
		} else {
			if left {
				adjacent_numbers = append(adjacent_numbers, parseNumbetAt(board, y-1, x-1))
			}

			if right {
				adjacent_numbers = append(adjacent_numbers, parseNumbetAt(board, y-1, x+1))
			}
		}
	}

	// d
	if x != 0 && strings.Contains(numbers, string((*board)[y][x-1])) {
		adjacent_numbers = append(adjacent_numbers, parseNumbetAt(board, y, x-1))
	}

	// e
	if x != (len((*board)[0])-1) && strings.Contains(numbers, string((*board)[y][x+1])) {
		adjacent_numbers = append(adjacent_numbers, parseNumbetAt(board, y, x+1))
	}

	if y != (len(*board) - 1) {
		left := false
		center := false
		right := false

		// e
		if x != 0 && strings.Contains(numbers, string((*board)[y+1][x-1])) {
			left = true
		}

		// f
		if strings.Contains(numbers, string((*board)[y+1][x])) {
			center = true
		}

		// g
		if x != (len((*board)[0])-1) && strings.Contains(numbers, string((*board)[y+1][x+1])) {
			right = true
		}

		if center {
			adjacent_numbers = append(adjacent_numbers, parseNumbetAt(board, y+1, x))
		} else {
			if left {
				adjacent_numbers = append(adjacent_numbers, parseNumbetAt(board, y+1, x-1))
			}

			if right {
				adjacent_numbers = append(adjacent_numbers, parseNumbetAt(board, y+1, x+1))
			}
		}
	}

	if len(adjacent_numbers) == 2 {
		return true, adjacent_numbers[0] * adjacent_numbers[1]
	}

	return false, 0
}

func SolveB() {

	fmt.Println("Day 3 Problem B")

	lines := utils.ReadLines("day3/input/b.input")

	total := 0

	for y, line := range lines {
		for x, char := range line {
			if string(char) == "*" {
				to_add, amount := isAdjacentToNumber(&lines, y, x)

				if to_add {
					total += amount
				}
			}
		}
	}

	fmt.Println(total)
}
