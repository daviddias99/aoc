package day5

import (
	"aoc/2023/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func SolveB() {

	fmt.Println("Day 5 Problem B")

	lines := utils.ReadLines("day5/input/b.input")

	seeds_line := lines[0]
	seeds_start_index := strings.Index(seeds_line, ":") + 2
	seeds_str := strings.Fields(seeds_line[seeds_start_index:])
	seeds := []int{}
	category_maps := [][]category_range{}

	for _, seed_str := range seeds_str {
		val, _ := strconv.Atoi(seed_str)
		seeds = append(seeds, val)
	}

	category_map := []category_range{}
	for i := 3; i < len(lines); i++ {

		if lines[i] == "" {
			category_maps = append(category_maps, category_map)
			category_map = []category_range{}
			i++
			continue
		}

		map_string := strings.Fields(lines[i])
		dest_low, _ := strconv.Atoi(map_string[0])
		src_low, _ := strconv.Atoi(map_string[1])
		range_size, _ := strconv.Atoi(map_string[2])

		new_range := category_range{
			src_lower_bound: src_low,
			src_upper_bound: src_low + range_size,
			dest_lower_bound: dest_low,
			dest_upper_bound: dest_low + range_size,
			size: range_size}

		category_map = append(category_map, new_range)
	}

	category_maps = append(category_maps, category_map)
	top := math.MaxInt

	for i := 0; i < len(seeds); i+=2 {
		fmt.Println(seeds[i], seeds[i] + seeds[i+1])
		for seed := seeds[i]; seed < seeds[i] + seeds[i+1]; seed++ {
			current_val := seed
			for _, category := range category_maps {			
				for _, category_range := range category {
					if (current_val >= category_range.src_lower_bound) && (current_val <= category_range.src_upper_bound) {
						current_val = category_range.dest_lower_bound + (current_val - category_range.src_lower_bound)
						break
					}				
				}
			}
			top = min(current_val, top)
		}
	}

	fmt.Println(top)
}
