package day12

import (
	"aoc/2023/utils"
	"fmt"
	"strconv"
	"strings"
)

// procedure backtrack(P, c) is
//     if reject(P, c) then return
//     if accept(P, c) then output(P, c)
//     s ← first(P, c)
//     while s ≠ NULL do
//         backtrack(P, s)
//         s ← next(P, s)

func SolveB() {

	fmt.Println("Day 12 Problem B")

	lines := utils.ReadLines("day12/input/b.input")

	total := 0

	for _, line := range lines {
		parts := strings.Split(line, " ")

		plist := parts[0] + "?" + parts[0] + "?" + parts[0] + "?" + parts[0] + "?" + parts[0]
		tlist := parts[1] + "," + parts[1] + "," + parts[1] + "," + parts[1] + "," + parts[1]

		solutions := []string{}

		targets_str := strings.Split(tlist, ",")
		targets := []int{}

		amount_needed := 0

		for _, target_str := range targets_str {
			val, _ := strconv.Atoi(target_str)
			targets = append(targets, val)
			amount_needed += val
		}

		FindSolution(plist, &targets, 0, &solutions, amount_needed)
		fmt.Println(len(solutions))
		total += len(solutions)
	}

	fmt.Println(total)
}
