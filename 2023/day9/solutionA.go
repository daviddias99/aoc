package day9

import (
	"aoc/2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func ConvertToSequence(str string) []int {
	result := []int{}
	fields := strings.Fields(str)

	for _, field := range fields {
		val, _ := strconv.Atoi(field)
		result = append(result, val)
	}

	return result
}

func ExtrapolateDown(sequence []int) [][]int {
	sequences := [][]int{sequence}

	for !utils.AllZeros(sequences[len(sequences)-1]) {
		current_sequence := sequences[len(sequences)-1]
		new_sequence := []int{}

		for i := 0; i < len(current_sequence)-1; i++ {
			new_sequence = append(new_sequence, current_sequence[i+1]-current_sequence[i])
		}
		sequences = append(sequences, new_sequence)
	}
	fmt.Println(sequences)

	return sequences
}

func ExtrapolateNextValue(sequence []int) int {
	sequences := ExtrapolateDown(sequence)

	for i := len(sequences) - 2; i >= 0; i-- {
		sequence_to_update := sequences[i]
		origin_sequence := sequences[i+1]

		new_value := origin_sequence[len(origin_sequence)-1] + sequence_to_update[len(origin_sequence)-1]
		sequence_to_update = append(sequence_to_update, new_value)

		sequences[i] = sequence_to_update
	}

	first_sequence := sequences[0]

	return first_sequence[len(first_sequence)-1]
}

func SolveA() {

	fmt.Println("Day 9 Problem A")

	lines := utils.ReadLines("day9/input/a.input")

	total := 0

	for _, line := range lines {
		starting_sequence := ConvertToSequence(line)
		total += ExtrapolateNextValue(starting_sequence)
	}

	fmt.Println(total)
}
