package day4

import (
	"aoc/2023/utils"
	"fmt"
	"math"
	"strings"
)

func parse_game(line string) int {
	game_map := make(map[string]int)
	games_start_index := strings.Index(line, ":") + 2
	plays := strings.Split(line[games_start_index:], "|")
	winning_numbers := strings.Fields(plays[0])
	play_numbers := strings.Fields(plays[1])

	total := 0.0
	
	for _, v := range winning_numbers {
		game_map[v] = 1
	}

	for _, v := range play_numbers {
		if game_map[v] == 1 {
			// fmt.Println("Value", v)
			total++
		}
	}


	// fmt.Println(winning_numbers)
	// fmt.Println(play_numbers)
	// fmt.Println("Total", total)

	return int(math.Pow(2, total - 1))
}


func SolveA() {

	fmt.Println("Day 4 Problem A")

	lines := utils.ReadLines("day4/input/a.input")
	total := 0

	for _, line := range lines {
		total += parse_game(line)
	}

	fmt.Println(total)
}
