package day10

import (
	"aoc/2023/utils"
	"errors"
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

func ChooseNextBasedOnTile(from position, to position, board *[]string) (position, error) {
	tile := BoardAt(board, to)
	if tile == '|' {
		// check for validity
		if from.x != to.x {
			return error_position, errors.New("invalid position")
		}

		if from.y < to.y {
			return position{x: to.x, y: to.y + 1}, nil
		} else {
			return position{x: to.x, y: to.y - 1}, nil
		}
	}

	if tile == '-' {

		if from.y != to.y {
			return error_position, errors.New("invalid position")
		}

		if from.x < to.x {
			return position{x: to.x + 1, y: to.y}, nil
		} else {
			return position{x: to.x - 1, y: to.y}, nil
		}
	}

	if tile == 'L' {

		if (from.y > to.y) || (from.x < to.x) {
			return error_position, errors.New("invalid position")
		}

		if from.y < to.y {
			return position{x: to.x + 1, y: to.y}, nil
		} else {
			return position{x: to.x, y: to.y - 1}, nil
		}
	}

	if tile == 'J' {
		if (from.y > to.y) || (from.x > to.x) {
			return error_position, errors.New("invalid position")
		}

		if from.y < to.y {
			return position{x: to.x - 1, y: to.y}, nil
		} else {
			return position{x: to.x, y: to.y - 1}, nil
		}
	}

	if tile == '7' {
		if (from.y < to.y) || (from.x > to.x) {
			return error_position, errors.New("invalid position")
		}

		if from.y > to.y {
			return position{x: to.x - 1, y: to.y}, nil
		} else {
			return position{x: to.x, y: to.y + 1}, nil
		}
	}

	if tile == 'F' {
		if (from.y < to.y) || (from.x < to.x) {
			return error_position, errors.New("invalid position")
		}

		if from.y > to.y {
			return position{x: to.x + 1, y: to.y}, nil
		} else {
			return position{x: to.x, y: to.y + 1}, nil
		}
	}

	if tile == 'S' {
		return position{x: 0, y: 0}, errors.New("invalid pos")
	}

	return position{x: 0, y: 0}, errors.New("invalid pos")
}

func SolveA() {

	fmt.Println("Day 10 Problem A")

	board := utils.ReadLines("day10/input/a.input")

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

	for BoardAt(&board, next_tile_pos) != 'S' {
		previous_next := next_tile_pos
		next_tile_pos, _ = ChooseNextBasedOnTile(current_tile_post, next_tile_pos, &board)
		current_tile_post = previous_next
		travelled++
	}

	fmt.Println(math.Ceil(float64(travelled) / 2))
}
