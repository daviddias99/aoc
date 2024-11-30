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

func CheckForPartialValidity(sequence string, targets []int) bool {
	target_index := 0
	current_val_size := 0

	for _, val := range sequence {

		if val == '?' {
			return true
		} else if val == '.' {

			if current_val_size > 0 {

				if current_val_size != targets[target_index] {
					return false
				}

				target_index++
				current_val_size = 0
			}

		} else if val == '#' {
			current_val_size++

			if (target_index == len(targets)) || (current_val_size > targets[target_index]) {
				return false
			}
		}
	}

	if current_val_size > 0 {

		if current_val_size != targets[target_index] {
			return false
		}
		target_index++
	}

	return target_index == len(targets)
}

func FindSolution(sequence string, targets *[]int, current_index int, solutions *[]string, amount_needed int) {

	fmt.Println(sequence)
	if !CheckForPartialValidity(sequence, *targets) {
		return
	}

	if current_index == len(sequence) {
		*solutions = append(*solutions, sequence)
		return
	}

	total := 0

	for _, v := range sequence {
		if v == '.' {
			continue
		} else {
			total++
		}
	}

	if total < amount_needed {
		return
	}

	current_val := sequence[current_index]

	if current_val == '?' {
		new_seq_dot := sequence[:current_index] + "." + sequence[current_index+1:]
		new_seq_hash := sequence[:current_index] + "#" + sequence[current_index+1:]

		FindSolution(new_seq_dot, targets, current_index+1, solutions, amount_needed)
		FindSolution(new_seq_hash, targets, current_index+1, solutions, amount_needed)
	} else {
		FindSolution(sequence, targets, current_index+1, solutions, amount_needed)
	}
}

func SolveA() {

	fmt.Println("Day 12 Problem A")

	lines := utils.ReadLines("day12/input/a_ex.input")

	total := 0

	for _, line := range lines {
		parts := strings.Split(line, " ")

		solutions := []string{}

		targets_str := strings.Split(parts[1], ",")
		targets := []int{}

		amount_needed := 0

		for _, target_str := range targets_str {
			val, _ := strconv.Atoi(target_str)
			targets = append(targets, val)
			amount_needed += val
		}

		FindSolution(parts[0], &targets, 0, &solutions, amount_needed)
		// fmt.Println(len(solutions))
		total += len(solutions)
	}

	fmt.Println(total)
}
