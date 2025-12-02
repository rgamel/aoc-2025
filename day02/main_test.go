package main

import "testing"

type test struct {
	in  int
	out bool
}

func TestInvalidId(t *testing.T) {

	tests := []test{
		{in: 11, out: true},
		{in: 22, out: true},
		{in: 99, out: true},
		{in: 1010, out: true},
		{in: 222222, out: true},
		{in: 446446, out: true},
		{in: 38593859, out: true},
		{in: 1, out: false},
		{in: 123, out: false},
		{in: 12123, out: false},
	}

	for _, tt := range tests {
		r := IsInvalidId(tt.in)
		if r != tt.out {
			t.Errorf("got %v, want %v", r, tt.out)
		}
	}
}

func TestAddInvalidId(t *testing.T) {
	type addtest struct {
		in  int
		out int
	}
	tests := []addtest{
		{in: 11, out: 11},
		{in: 12, out: 0},
		{in: 22, out: 22},
		{in: 23, out: 0},
	}

	for _, tt := range tests {
		r := AddInvalidId(tt.in, 0)
		if r != tt.out {
			t.Errorf("got %d, want %d", r, tt.out)
		}
	}
}

func TestParseRange(t *testing.T) {
	type rangetest struct {
		in    string
		start int
		end   int
	}

	tests := []rangetest{
		{in: "1-2", start: 1, end: 2},
		{in: "999-1001", start: 999, end: 1001},
	}

	for _, tt := range tests {
		rStart, rEnd := ParseRange(tt.in)

		if rStart != tt.start {
			t.Errorf("got %d, want %d", rStart, tt.start)
		}

		if rEnd != tt.end {
			t.Errorf("got %d, want %d", rEnd, tt.end)
		}
	}
}
