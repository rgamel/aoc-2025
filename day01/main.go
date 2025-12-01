package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	lines := []string{}
	curr := 50
	count := 0

	// read input file
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("Problem reading file: %v", err)
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		next := s.Text()
		lines = append(lines, next)
	}

	for _, line := range lines {
		dir, amt := ParseDirection(line)

		rotFn := RotateRight

		if dir == "L" {
			rotFn = RotateLeft
		}
		curr = rotFn(curr, amt)
		CheckForZero(curr, &count)
	}

	fmt.Printf("password: %d\n", count)
}

func ParseDirection(input string) (dir string, amt int) {
	dir = string(input[0])

	amt, err := strconv.Atoi((input[1:]))
	if err != nil {
		log.Fatalf("could not read line: %s", input)
	}

	return dir, amt
}

func RotateRight(start int, dist int) (end int) {
	end = start + dist
	if end > 99 {
		end = end % 100
	}
	return end
}

func RotateLeft(start int, dist int) (end int) {
	end = start - dist
	for end < 0 {
		end = 100 + end
	}
	return end
}

func CheckForZero(input int, count *int) {
	if input == 0 {
		*count++
	}
}
