package main

import (
	"testing"

	utils "aoc2025/shared"
)

func TestPartA(t *testing.T) {
	lines := utils.ReadInput("./testinput.txt")
	ranges, ids := getInputs(lines)
	res := PartA(ids, ranges)

	if res != 3 {
		t.Errorf("Got %d want 3\n", res)
	}

}

func TestPartB(t *testing.T) {
	lines := utils.ReadInput("./testinput.txt")
	ranges, ids := getInputs(lines)
	res := PartB(ids, ranges)

	if res != 14 {
		t.Errorf("Got %d want 14\n", res)
	}
}

func TestIsInRange(t *testing.T) {
	tests := []struct {
		rangeStr IdRange
		id       int
		out      bool
	}{
		{IdRange{Min: 12340, Max: 12349}, 12340, true},
		{IdRange{Min: 12340, Max: 12349}, 12349, true},
		{IdRange{Min: 12340, Max: 12349}, 12339, false},
		{IdRange{Min: 12340, Max: 12349}, 12350, false},
	}

	for _, tt := range tests {
		res := tt.rangeStr.Has(tt.id)
		if tt.out != res {
			t.Errorf("got %v, want %v", res, tt.out)
		}
	}
}

func TestMergeSortedRanges(t *testing.T) {
	tests := []struct {
		ranges []IdRange
		out    int
	}{
		{ranges: []IdRange{
			{Min: 5, Max: 11},
			{Min: 1, Max: 10},
			{Min: 12, Max: 15},
		}, out: 2},
	}

	for _, tt := range tests {
		res := mergeSortedRanges(tt.ranges)
		t.Log(res)

		if len(res) != tt.out {
			t.Errorf("wanted %d, got %d\n", tt.out, len(res))
		}
	}

}
