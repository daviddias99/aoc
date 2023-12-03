package main


import (
	"aoc/2023/day1"
	"aoc/2023/day2"
	"aoc/2023/day3"
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
	}
}
