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
		log.Fatal("problem opening input file")
	}
	defer f.Close()

	joltage := 0
	maxJoltage := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		joltage += FindLineJoltage(text)
		maxJoltage += FindMaxLineJoltage(text)
	}
	fmt.Printf("part a: %d\n", joltage)
	fmt.Printf("part b: %d\n", maxJoltage)
}

func FindLineJoltage(line string) int {
	result := 0
	for i := range line {
		// TODO: exit early if 9
		for j := i + 1; j < len(line); j++ {
			// TODO: exit early if 9
			v, err := strconv.Atoi(fmt.Sprintf("%c%c", line[i], line[j]))
			if err != nil {
				log.Fatal("problem reading line")
			}
			if v > result {
				result = v
			}
		}
	}
	return result
}

func FindMaxLineJoltage(line string) int {
	p1 := 0
	p2 := len(line) - 12
	output := ""

	for p2 < len(line) {
		subStr := line[p1 : p2+1]
		subMax, idx, _ := findMaxDigit(subStr)
		output = output + fmt.Sprint(subMax)
		p1 += idx + 1
		p2++
	}
	parsed, _ := strconv.Atoi(output)
	return parsed
}

func findMaxStartIndex(line string, batteryCount int) int {
	return len(line) - batteryCount
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
