package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		fmt.Println(line)
	}
}

func PartA() {}
func PartB() {}
