package main

import (
	"aoc/2023/day1"
	"aoc/2023/day10"
	"aoc/2023/day11"
	"aoc/2023/day12"
	"aoc/2023/day2"
	"aoc/2023/day3"
	"aoc/2023/day4"
	"aoc/2023/day5"
	"aoc/2023/day6"
	"aoc/2023/day7"
	"aoc/2023/day8"
	"aoc/2023/day9"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Advent of Code: 2023")

	id := os.Args[1]
	
	switch id {
	case "1":
		day1.SolveA()
		day1.SolveB()
	case "2":
		day2.SolveA()
		day2.SolveB()
	case "3":
		day3.SolveA()
		day3.SolveB()
	case "4":
		day4.SolveA()
		day4.SolveB()
	case "5":
		day5.SolveA()
		day5.SolveB()
	case "6":
		day6.SolveA()
		day6.SolveB()
	case "7":
		day7.SolveA()
		day7.SolveB()
	case "8":
		day8.SolveA()
		day8.SolveB()
	case "9":
		day9.SolveA()
		day9.SolveB()
	case "10":
		day10.SolveA()
		day10.SolveB()
	case "11":
		day11.SolveA()
		day11.SolveB()
	case "12":
		day12.SolveA()
		day12.SolveB()
	}
}
