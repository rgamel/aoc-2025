package main

import "testing"

type Test struct {
	in  int
	out bool
}

func TestInvalidId(t *testing.T) {

	tests := []Test{
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
		{in: 22, out: 22},
	}

	for _, tt := range tests {
		r := AddInvalidId(tt.in, 0, func(id int) bool { return true })
		if r != tt.out {
			t.Errorf("got %d, want %d", r, tt.out)
		}
	}

	validtests := []addtest{
		{in: 12, out: 0},
		{in: 23, out: 0},
	}

	for _, tt := range validtests {
		r := AddInvalidId(tt.in, 0, func(id int) bool { return false })
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
		r := ParseRange(tt.in)

		if r.Min != tt.start {
			t.Errorf("got %d, want %d", r.Min, tt.start)
		}

		if r.Max != tt.end {
			t.Errorf("got %d, want %d", r.Max, tt.end)
		}
	}
}

func TestIsInvalidIdB(t *testing.T) {
	tests := []Test{
		{in: 11, out: true},
		{in: 12, out: false},
		{in: 22, out: true},
		{in: 23, out: false},
		{in: 99, out: true},
		{in: 98, out: false},
		{in: 111, out: true},
		{in: 112, out: false},
		{in: 999, out: true},
		{in: 998, out: false},
		{in: 1010, out: true},
		{in: 10100, out: false},
		{in: 1188511885, out: true},
	}

	for _, tt := range tests {
		r := IsInvalidIdB(tt.in)
		if r != tt.out {
			t.Errorf("\nin: %d\ngot %v, want %v", tt.in, r, tt.out)
		}
	}

}
