package day11

import (
	"aoc/2023/utils"
	"fmt"
	"math"
)

type position struct {
	x int
	y int
}

var error_position = position{x: -1, y: -1}

func BoardAt(board *[]string, pos position) byte {
	return (*board)[pos.y][pos.x]
}

func SolveA() {

	fmt.Println("Day 11 Problem A")

	board := utils.ReadLines("day11/input/a_ex.input")

	column_tracker := []bool{}
	galaxies := []position{}
	empty_rows := []int{}
	empty_row := ""

	for i := 0; i < len(board[0]); i++ {
		column_tracker = append(column_tracker, false)
		empty_row = empty_row + "."
	}

	for y, row := range board {
		galaxy_found := false
		for x, space := range row {
			pos := position{x: x, y: y}

			if space == '#' {
				galaxies = append(galaxies, pos)
				column_tracker[x] = true
				galaxy_found = true
			}
		}

		if !galaxy_found {
			empty_rows = append(empty_rows, y)
		}
	}

	galaxies_new := []position{}

	for _, galaxy := range galaxies {
		galaxies_new = append(galaxies_new, galaxy)
	}

	for _, empty_row := range empty_rows {
		for i, galaxy := range galaxies {
			if galaxy.y > empty_row {
				galaxies_to_update := galaxies_new[i]
				galaxies_to_update.y++
				galaxies_new[i] = galaxies_to_update
			}
		}
	}
	for x, filled_column := range column_tracker {
		if !filled_column {
			for i, galaxy := range galaxies {

				if galaxy.x > x {
					galaxies_to_update := galaxies_new[i]
					galaxies_to_update.x++
					galaxies_new[i] = galaxies_to_update
				}
			}
		}
	}

	total := 0

	for i := 0; i < len(galaxies_new)-1; i++ {
		for j := i + 1; j < len(galaxies_new); j++ {
			galaxy_a := galaxies_new[i]
			galaxy_b := galaxies_new[j]
			space_between := int(math.Abs(float64(galaxy_b.y-galaxy_a.y))) + int(math.Abs(float64(galaxy_b.x-galaxy_a.x)))
			fmt.Println("Space between galaxy", i+1, galaxy_a, "and", j+1, galaxy_b, "is", space_between)
			total += space_between
		}
	}

	fmt.Println(total)
}
