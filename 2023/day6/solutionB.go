package day6

import (
	"aoc/2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func SolveB() {
	fmt.Println("Day 6 Problem B")

	lines := utils.ReadLines("day6/input/b.input")

	times_start_index := strings.Index(lines[0], ":") + 1
	distances_start_index := strings.Index(lines[1], ":") + 1

	times_str := lines[0][times_start_index:]
	distances_str := lines[1][distances_start_index:]

	time, _ := strconv.Atoi(strings.ReplaceAll(times_str, " ", ""))
	distance, _ := strconv.Atoi(strings.ReplaceAll(distances_str, " ", ""))

	sol1, sol2 := SolveQuadraticAprox(-1, time, -distance)
	fmt.Println((sol2 - sol1 + 1))
}
