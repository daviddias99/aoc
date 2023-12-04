package day4

import (
	"aoc/2023/utils"
	"fmt"
	// "math"
	// "math"
	"strings"
)

func parse_game_b(line string) int {
	game_map := make(map[string]int)
	games_start_index := strings.Index(line, ":") + 2
	plays := strings.Split(line[games_start_index:], "|")
	winning_numbers := strings.Fields(plays[0])
	play_numbers := strings.Fields(plays[1])

	total := 0

	for _, v := range winning_numbers {
		game_map[v] = 1
	}

	for _, v := range play_numbers {
		if game_map[v] == 1 {
			total++
		}
	}

	return total 
}


func SolveB() {

	fmt.Println("Day 4 Problem B")

	lines := utils.ReadLines("day4/input/a.input")
	// total := 0

	totals := make([]int, len(lines) + 1)
	for i:= 1; i < len(totals); i++{
		totals[i] = 1
	}

	for i, total := range totals {
		if i == 0 {
			continue
		}

		game_return := parse_game_b(lines[i - 1])

		for j := i + 1; j < min(len(totals), i + game_return + 1) ; j++ {
			totals[j] += total			
		}
	}

	total := 0

	for _, v := range totals {
		total += v
	}

	fmt.Println(totals)
	fmt.Println(total)
}
