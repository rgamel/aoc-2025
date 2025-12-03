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

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		joltage += FindLineJoltage(scanner.Text())
	}

	fmt.Printf("part a: %d\n", joltage)
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
