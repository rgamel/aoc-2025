package main

import "testing"

type Test struct {
	in  string
	out int
}

func TestFindLineJoltage(t *testing.T) {
	tests := []Test{
		{in: "987654321111111", out: 98},
		{in: "811111111111119", out: 89},
		{in: "234234234234278", out: 78},
		{in: "818181911112111", out: 92},
	}

	for _, tt := range tests {
		res := FindLineJoltage(tt.in)
		if res != tt.out {
			t.Errorf("got %d, want %d", res, tt.out)
		}
	}

}

func TestFindMaxLineJoltage(t *testing.T) {
	tests := []Test{
		{in: "987654321111111", out: 987654321111},
		{in: "811111111111119", out: 811111111119},
		{in: "234234234234278", out: 434234234278},
		{in: "818181911112111", out: 888911112111},
	}

	for _, tt := range tests {
		res := FindMaxLineJoltage(tt.in)
		if res != tt.out {
			t.Errorf("got %d, want %d", res, tt.out)
		}
	}

}

func TestFindMaxDigit(t *testing.T) {
	type MaxTest struct {
		in  string
		idx int
		out int
	}
	tests := []MaxTest{
		{in: "987654321", idx: 0, out: 9},
		{in: "89654321", idx: 1, out: 9},
		{in: "12344321", idx: 3, out: 4},
	}

	for _, tt := range tests {
		max, i, _ := findMaxDigit(tt.in)
		if max != tt.out {
			t.Errorf("wrong max, got %d, want %d", max, tt.out)
		}
		if i != tt.idx {
			t.Errorf("wrong idx, got %d, want %d", i, tt.idx)
		}
	}

}
