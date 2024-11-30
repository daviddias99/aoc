package day10

import (
	"aoc/2023/utils"
	"errors"
	"fmt"
)

func IsInBorder(border_positions *[]position, pos_to_check position) bool {
	for _, border_pos := range *border_positions {
		if border_pos == pos_to_check {
			return true
		}
	}

	return false
}

func FilterBorders(board *[]string, border_positions *[]position) []position {
	filtered_positions := []position{}
	for _, border_pos := range *border_positions {
		tile := (*board)[border_pos.y][border_pos.x]
		if tile == 'F' || tile == '7' || tile == '-' || tile == 'S' {
			continue
		} else {
			filtered_positions = append(filtered_positions, border_pos)
		}
	}

	return filtered_positions
}

func CountBorderPositionsToRight(border_positions *[]position, pos_to_check position, max_width int) int {
	result := 0
	for i := pos_to_check.x + 1; i < max_width; i++ {
		pos := position{x: i, y: pos_to_check.y}
		is_in_border := IsInBorder(border_positions, pos)

		if is_in_border{
			result++
		} 
	}

	return result
}

func SolveB() {

	fmt.Println("Day 10 Problem B")

	board := utils.ReadLines("day10/input/b_ex.input")

	var start_pos position

	for y, line := range board {
		for x, tile := range line {
			if tile == 'S' {
				start_pos = position{x: x, y: y}
			}
		}
	}

	var next_tile_pos position
	var err error = errors.New("no valid starting position found")

	if (start_pos.y != 0) && err != nil {
		_, err = ChooseNextBasedOnTile(start_pos, position{x: start_pos.x, y: start_pos.y - 1}, &board)
		next_tile_pos = position{x: start_pos.x, y: start_pos.y - 1}
	}

	if (start_pos.x != 0) && err != nil {
		_, err = ChooseNextBasedOnTile(start_pos, position{x: start_pos.x - 1, y: start_pos.y}, &board)
		next_tile_pos = position{x: start_pos.x - 1, y: start_pos.y}
	}

	if (start_pos.y != len(board)-1) && err != nil {
		_, err = ChooseNextBasedOnTile(start_pos, position{x: start_pos.x, y: start_pos.y + 1}, &board)
		next_tile_pos = position{x: start_pos.x, y: start_pos.y + 1}
	}

	if (start_pos.x != len(board)-1) && err != nil {
		_, err = ChooseNextBasedOnTile(start_pos, position{x: start_pos.x + 1, y: start_pos.y}, &board)
		next_tile_pos = position{x: start_pos.x + 1, y: start_pos.y}
	}

	if err != nil {
		panic(err)
	}

	current_tile_post := start_pos
	travelled := 0
	border_positions := []position{start_pos}
	insiders := 0

	for BoardAt(&board, next_tile_pos) != 'S' {
		previous_next := next_tile_pos
		next_tile_pos, _ = ChooseNextBasedOnTile(current_tile_post, next_tile_pos, &board)
		current_tile_post = previous_next
		travelled++
		border_positions = append(border_positions, current_tile_post)
	}

	filtered_borders := FilterBorders(&board, &border_positions)

	for y, line := range board {
		for x := range line {
			pos := position{x: x, y: y}
			is_border := false

			for _, border_pos := range border_positions {
				if border_pos == pos {
					is_border = true
				}
			}

			if is_border {
				continue
			}

			crossings := CountBorderPositionsToRight(&filtered_borders, pos, len(line))

			if crossings % 2 != 0 {
				insiders++
			}
		}
	}


	fmt.Println(insiders)
}
