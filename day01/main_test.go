package main

import "testing"

type test struct {
	in   int
	dist int
	out  int
}

func TestRotateRight(t *testing.T) {
	tests := []test{
		{in: 99, dist: 1, out: 0},
		{in: 0, dist: 1, out: 1},
		{in: 50, dist: 50, out: 0},
		{in: 51, dist: 30, out: 81},
	}
	for _, tt := range tests {
		result := RotateRight(tt.in, tt.dist)
		if result != tt.out {
			t.Errorf("wanted %d, got %d", tt.out, result)
		}
	}
}

func TestRotateLeft(t *testing.T) {
	tests := []test{
		{in: 99, dist: 1, out: 98},
		{in: 0, dist: 1, out: 99},
		{in: 50, dist: 50, out: 0},
		{in: 51, dist: 30, out: 21},
		{in: 20, dist: 100, out: 20},
	}
	for _, tt := range tests {
		result := RotateLeft(tt.in, tt.dist)
		if result != tt.out {
			t.Errorf("wanted %d, got %d", tt.out, result)
		}
	}
}

func TestCheckForZero(t *testing.T) {
	store := Store{count: 0}
	CheckForZero(1, store)

	if store.count != 0 {
		t.Errorf("wanted %d, got %d", 0, store.count)
	}

}

func TestParseDirection(t *testing.T) {
	input := "L25"

	dir, amt := ParseDirection(input)

	if dir != "L" {
		t.Errorf("Going the wrong way")
	}

	if amt != 25 {
		t.Errorf("Wrong amt")
	}
}
