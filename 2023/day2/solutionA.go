package day2

import (
	"aoc/2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func getMaxCubes(line string) (int, map[string]int) {
	game_map := make(map[string]int)
	game_id, err := strconv.Atoi(string(game_id_regex.FindStringSubmatch(line)[1]))

	if err != nil {
		panic("Could not parse game id")
	}

	games_start_index := strings.Index(line, ":") + 2
	plays := color_quantity_regex.FindAllStringSubmatch(line[games_start_index:], -1)

	for _, play := range plays {
		cube_count, err := strconv.Atoi(play[1])

		if err != nil {
			panic("Could not parse play")
		}

		game_map[play[2]] = max(cube_count, game_map[play[2]])
	}

	return game_id, game_map
}

func SolveA() {

	fmt.Println("Day 2 Problem A")

	lines := utils.ReadLines("day2/input/a.input")

	total := 0
	MAX_RED := 12
	MAX_GREEN := 13
	MAX_BLUE := 14

	for _, v := range lines {
		game_id, game_map := getMaxCubes(v)

		if (game_map["red"] > MAX_RED) || (game_map["blue"] > MAX_BLUE) || (game_map["green"] > MAX_GREEN) {
			continue
		}

		total += game_id
	}

	fmt.Println(total)
}
