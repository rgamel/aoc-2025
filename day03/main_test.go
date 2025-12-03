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
