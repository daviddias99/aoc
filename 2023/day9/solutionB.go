package day9

import (
	"aoc/2023/utils"
	"fmt"
)


func ExtrapolatePreviousValue(sequence []int) int {
	sequences := ExtrapolateDown(sequence)

	for i := len(sequences) - 2; i >= 0; i-- {
		sequence_to_update := sequences[i]
		origin_sequence := sequences[i + 1]

		new_value :=  sequence_to_update[0] - origin_sequence[0]
		sequence_to_update = append([]int{new_value}, sequence_to_update...)

		sequences[i] = sequence_to_update
	}

	first_sequence := sequences[0]

	return first_sequence[0]
}

func SolveB() {

	fmt.Println("Day 9 Problem B")

	lines := utils.ReadLines("day9/input/b.input")

	fmt.Println(len(lines))

	total := 0

	for _, line := range lines {
		starting_sequence := ConvertToSequence(line)
		total += ExtrapolatePreviousValue(starting_sequence)
	}

	fmt.Println(total)
}
