package day3

import (
	"aoc/2023/utils"
	"fmt"
	"strconv"
	"strings"
)

var symbols = "!\"#$%&'()*+,-/:;<=>?@[]^_{|}~"

func isAnyAdjacentSymbols(board *[]string, y int, x int) bool {

	// | a b c |
	// | d X e |
	// | f g h |

	if y != 0 {
		// a
		if x != 0 && strings.Contains(symbols, string((*board)[y-1][x-1])) {
			return true
		}

		// b
		if strings.Contains(symbols, string((*board)[y-1][x])) {
			return true
		}

		// c
		if x != (len((*board)[0])-1) && strings.Contains(symbols, string((*board)[y-1][x+1])) {
			return true
		}
	}

	// d
	if x != 0 && strings.Contains(symbols, string((*board)[y][x-1])) {
		return true
	}

	// e
	if x != (len((*board)[0])-1) && strings.Contains(symbols, string((*board)[y][x+1])) {
		return true
	}

	if y != (len(*board) - 1) {
		// f
		if x != 0 && strings.Contains(symbols, string((*board)[y+1][x-1])) {
			return true
		}

		// g
		if strings.Contains(symbols, string((*board)[y+1][x])) {
			return true
		}

		// h
		if x != (len((*board)[0])-1) && strings.Contains(symbols, string((*board)[y+1][x+1])) {
			return true
		}
	}

	return false
}

func SolveA() {

	fmt.Println("Day 3 Problem A")

	lines := utils.ReadLines("day3/input/a.input")

	total := 0

	for y, line := range lines {
		current_number := ""
		to_add := false
		for x, char := range line {

			is_numeric := char >= '0' && char <= '9'

			if  is_numeric {
				current_number += string(char)

				if isAnyAdjacentSymbols(&lines, y, x) {
					to_add = true
				}
			} 
			
			if (!is_numeric && (current_number != "")) || (is_numeric && (x == len(line)-1)) {
				value, err := strconv.Atoi(current_number)

				if err != nil {
					panic("Could not convert number")
				}

				if to_add {
					total += value
				}
				to_add = false
				current_number = ""
			}
		}
	}

	fmt.Println(total)
}
