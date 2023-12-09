package day8

import (
	"aoc/2023/utils"
	"fmt"
	"strings"
)

func AllTrue(to_test []bool) bool {
	for _, v := range to_test {
		if !v {
			return false
		}
	}

	return true
}

func SolveB() {

	fmt.Println("Day 8 Problem B")

	lines := utils.ReadLines("day8/input/b.input")

	directions_str := lines[0]
	directions := DirectionsFromString(directions_str)
	dir_map := ParseInstructionsIntoMap(lines[2:])

	currents := []string{}
	counters := []int{}

	for key := range dir_map {
		if strings.HasSuffix(key, "A") {
			currents = append(currents, key)
		}
	}

	for key := range currents {
		var i int

		for i = 0; !strings.HasSuffix(currents[key], "Z"); i++ {
			currents[key] = dir_map[currents[key]][directions[i % len(directions)]]
		}

		counters = append(counters, i)
	}

	fmt.Println(utils.LCM(counters...))
}
