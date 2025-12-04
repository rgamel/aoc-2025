package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("problem reading input file: %v\n", err)
	}
	defer f.Close()

	lines := []string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	run(lines, "Part A", PartA)
	run(lines, "Part B", PartB)
}

func run(lines []string, label string, cb func(lines []string) int) {
	start := time.Now()
	res := cb(lines)
	end := time.Now()
	diff := end.Sub(start)
	fmt.Printf("%s: %d, duration: %d\n", label, res, diff.Microseconds())
}

func PartA(matrix []string) int {
	rolls := 0
	for rowI := range matrix {
		for colI := range matrix {
			rolls += ScanAdjacent(rowI, colI, matrix)
		}
	}
	return rolls
}

func PartB(matrix []string) int {
	rolls := 0
	next := matrix
	for {
		curr := 0
		for row := 0; row < len(matrix); row++ {
			for col := 0; col < len(matrix[row]); col++ {
				res := ScanAdjacent(row, col, matrix)
				if res == 1 {
					rolls++
					curr++
					next[row] = matrix[row][:col] + "." + matrix[row][col+1:]
				}
			}
		}
		if curr == 0 {
			break
		}
		matrix = next
	}
	return rolls
}

func ScanAdjacent(row int, col int, matrix []string) int {
	if matrix[row][col] == '.' {
		return 0
	}
	adjRolls := 0
	for direction := range 9 {
		if direction == 4 {
			// this is self
			continue
		}
		n_row := row + ((direction % 3) - 1)
		n_col := col + ((direction / 3) - 1)
		if n_row >= 0 && n_row < len(matrix) && n_col >= 0 && n_col < len(matrix[0]) {
			// a valid neighbor
			if matrix[n_row][n_col] == '@' {
				adjRolls++
			}
			if adjRolls > 3 {
				return 0
			}
		}
	}
	return 1
}
