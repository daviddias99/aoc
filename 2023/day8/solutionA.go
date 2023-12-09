package day8

import (
	"aoc/2023/utils"
	"fmt"
)

func DirectionsFromString(str string) []int {
	directions := []int{}

	for _, direction := range str {

		if direction == 'R' {
			directions = append(directions, 1)
		} else {
			directions = append(directions, 0)
		}
	}

	return directions
}

func ParseInstructionsIntoMap(instructions []string) map[string][2]string {
	dir_map := make(map[string][2]string)

	for _, line := range instructions {
		dir_map[line[:3]] = [2]string{line[7:10], line[12:15]}
	}

	return dir_map
}

func SolveA() {

	fmt.Println("Day 8 Problem A")

	lines := utils.ReadLines("day8/input/a.input")

	directions_str := lines[0]
	directions := DirectionsFromString(directions_str)
	dir_map := ParseInstructionsIntoMap(lines[2:])

	current := "AAA"
	var total_steps int

	for total_steps = 0; current != "ZZZ"; total_steps++{
		current = dir_map[current][directions[total_steps % len(directions)]]
	}

	fmt.Println(total_steps)
}
