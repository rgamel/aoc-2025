package main

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func getTestMatrix() []string {
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

func TestScanAdjacent(t *testing.T) {
	tests := []struct {
		row int
		col int
		res int
	}{
		{0, 0, 0},
		{0, 1, 0},
		{0, 2, 1},
		{0, 3, 1},
		{0, 4, 0},
		{0, 5, 1},
		{0, 6, 1},
		{0, 7, 0},
		{0, 8, 1},
		{0, 9, 0},
		{1, 0, 1},
		{1, 1, 0},
		{1, 1, 0},
		{9, 8, 1},
	}
	matrix := getTestMatrix()

	for _, tt := range tests {
		res := CheckNeighbors(tt.row, tt.col, matrix)
		if tt.res != res {
			t.Errorf("got %d want %d", res, tt.res)

		}
	}
}

func TestPartA(t *testing.T) {
	matrix := getTestMatrix()
	res := PartA(matrix)

	if res != 13 {
		t.Errorf("Got %d, want 13", res)
	}
}

func TestPartB(t *testing.T) {
	matrix := getTestMatrix()
	res := PartB(matrix)

	if res != 43 {
		t.Errorf("Got %d, want 43", res)
	}
}
