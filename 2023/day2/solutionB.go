package day2

import (
	"aoc/2023/utils"
	"fmt"
)

func SolveB() {

	fmt.Println("Day 2 Problem B")

	lines := utils.ReadLines("day2/input/b.input")

	total := 0

	for _, v := range lines {
		_, game_map := getMaxCubes(v)

		total += game_map["red"] * game_map["green"] * game_map["blue"]
	}

	fmt.Println(total)
}
