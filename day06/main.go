package main

import (
	"fmt"
	"time"

	utils "aoc2025/shared"
)

func main() {
	lines := utils.ReadInput("./input.txt")
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
	return 0
}

func PartB(matrix []string) int {
	return 0
}
