package main

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func getTestInput() []string {
	f, err := os.Open("./testinput.txt")
	if err != nil {
		log.Fatalf("problem opening test file: %v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	matrix := []string{}
	for scanner.Scan() {
		matrix = append(matrix, scanner.Text())
	}
	return matrix
}

func TestPartA(t *testing.T) {
	// input := getTestInput()
	// res := PartA(input)
	t.Errorf("not implemented")
}

func TestPartB(t *testing.T) {
	// input := getTestInput()
	// res := PartB(input)
	t.Errorf("not implemented")
}
