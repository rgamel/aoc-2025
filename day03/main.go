package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("problem opening input file: %v", err)
	}
	defer f.Close()

	joltage := 0
	maxJoltage := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		lineJoltage, err := FindMaxLineJoltage(text, 2)
		if err != nil {
			log.Fatal(err)
		}
		joltage += lineJoltage
		maxLineJoltage, err := FindMaxLineJoltage(text, 12)
		maxJoltage += maxLineJoltage
	}
	fmt.Printf("part a: %d\n", joltage)
	fmt.Printf("part b: %d\n", maxJoltage)
}

func FindMaxLineJoltage(line string, batteries int) (int, error) {
	p1 := 0
	p2 := len(line) - batteries
	output := ""

	for p2 < len(line) {
		subStr := line[p1 : p2+1]
		subMax, idx, err := findMaxDigit(subStr)
		if err != nil {
			return -1, err
		}
		output = output + fmt.Sprint(subMax)
		p1 += idx + 1
		p2++
	}

	parsed, err := strconv.Atoi(output)
	if err != nil {
		log.Fatalf("problem converting output: %v", err)
	}

	return parsed, nil
}

func findMaxDigit(s string) (int, int, error) {
	maxDigit := 0
	idx := -1
	for i, char := range s {
		if char >= '0' && char <= '9' {
			digit, err := strconv.Atoi(string(char))
			if err != nil {
				return -1, -1, fmt.Errorf("error converting char to digit %w", err)
			}

			if digit > maxDigit {
				maxDigit = digit
				idx = i
			}
		}
	}

	return maxDigit, idx, nil
}
