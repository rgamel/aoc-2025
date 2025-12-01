package main

import "testing"

type test struct {
	in    int
	dist  int
	out   int
	count int
}

func TestRotateRight(t *testing.T) {

	tests := []test{
		{in: 99, dist: 1, out: 0, count: 1},
		{in: 0, dist: 1, out: 1, count: 0},
		{in: 50, dist: 50, out: 0, count: 1},
		{in: 51, dist: 30, out: 81, count: 0},
		{in: 50, dist: 300, out: 50, count: 3},
		{in: 50, dist: 150, out: 0, count: 2},
	}
	for _, tt := range tests {
		store := Store{Count: 0}
		result := RotateRight(tt.in, tt.dist, &store)
		if result != tt.out {
			t.Errorf("wanted %d, got %d", tt.out, result)
		}
		if store.Count != tt.count {
			t.Errorf("wanted %d, got %d", tt.count, store.Count)
		}
	}
}

func TestRotateLeft(t *testing.T) {
	tests := []test{
		{in: 99, dist: 1, out: 98, count: 0},
		{in: 2, dist: 1, out: 1, count: 0},
		{in: 0, dist: 1, out: 99, count: 0},
		{in: 50, dist: 50, out: 0, count: 1},
		{in: 51, dist: 30, out: 21, count: 0},
		{in: 20, dist: 100, out: 20, count: 1},
		{in: 20, dist: 200, out: 20, count: 2},
		{in: 50, dist: 150, out: 0, count: 2},
	}
	for _, tt := range tests {
		store := Store{Count: 0}
		result := RotateLeft(tt.in, tt.dist, &store)
		if result != tt.out {
			t.Errorf("wanted %d, got %d", tt.out, result)
		}
		if store.Count != tt.count {
			t.Errorf("wanted %d, got %d", tt.count, store.Count)
		}
	}
}

func TestCheckForZero(t *testing.T) {
	store := Store{Count: 0}
	CheckForZero(1, &store)

	if store.Count != 0 {
		t.Errorf("wanted %d, got %d", 0, store.Count)
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
