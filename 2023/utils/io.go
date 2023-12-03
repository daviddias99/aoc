package utils

import (
	"bufio"
	"os"
)

func ReadLines(path string) []string {
	file, err := os.Open(path)

	if err != nil {
		panic("Could not read input")
	}

	// defer makes the function be called at the end of the block
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		panic("Could not read input")
	}

	return lines
}
