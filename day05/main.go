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

func PartA(lines []string) int {
	return 0
}

func PartB(lines []string) int {
	return 0
}
