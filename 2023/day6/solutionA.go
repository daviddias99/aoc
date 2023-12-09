package day6

import (
	"aoc/2023/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func SolveQuadraticAprox(a int, b int, c int) (int, int) {

	sqrt_val := math.Sqrt(float64(b * b - 4 * a * c))
	sol1 := (-float64(b) - sqrt_val) / float64(2 * a)
	sol2 := (-float64(b) + sqrt_val) / float64(2 * a)

	return int(math.Ceil(sol2 + 0.001)), int(math.Floor(sol1 - 0.001))
}

func SolveA() {

	fmt.Println("Day 6 Problem A")

	lines := utils.ReadLines("day6/input/a.input")

	times_start_index := strings.Index(lines[0], ":") + 1
	distances_start_index := strings.Index(lines[1], ":") + 1

	times_str := lines[0][times_start_index:]
	distances_str := lines[1][distances_start_index:]

	times_str_arr := strings.Fields(times_str)
	distances_str_arr := strings.Fields(distances_str)

	times_arr := []int{}
	distances_arr := []int{}

	for _, v := range times_str_arr {
		value, _ := strconv.Atoi(v)
		times_arr = append(times_arr, value)
	}

	for _, v := range distances_str_arr {
		value, _ := strconv.Atoi(v)
		distances_arr = append(distances_arr, value)
	}

	total := 1

	for i := 0; i < len(times_arr); i++ {
		sol1, sol2 := SolveQuadraticAprox(-1, times_arr[i], -distances_arr[i])
		total *= (sol2 - sol1 + 1)
	}

	// fmt.Println(times_arr)
	// fmt.Println(distances_arr)
	fmt.Println(total)
}
