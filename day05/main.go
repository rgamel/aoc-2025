package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
	"time"

	utils "aoc2025/shared"
)

type IdRange struct {
	Min, Max int
}

func (r *IdRange) Has(id int) bool {
	return r.Min <= id && r.Max >= id
}

func (r *IdRange) Overlaps(other IdRange) bool {
	if r.Min == other.Min {
		return true
	}
	if r.Min < other.Min {
		return other.Min <= r.Max
	}

	return other.Max >= r.Min

}

func (r *IdRange) Merge(other IdRange) {
	r.Min = min(r.Min, other.Min)
	r.Max = max(r.Max, other.Max)
}

type Set struct {
	elements map[int]struct{}
}

func NewSet() *Set {
	return &Set{
		elements: make(map[int]struct{}),
	}
}

func (s *Set) Add(value int) {
	s.elements[value] = struct{}{}
}

func (s *Set) Size() int {
	return len(s.elements)
}

func main() {
	lines := utils.ReadInput("./input.txt")
	ranges, ids := getInputs(lines)
	run(ids, ranges, "Part A", PartA)
	run(ids, ranges, "Part B", PartB)
}

func run(ids []int, ranges []IdRange, label string, cb func(ids []int, ranges []IdRange) int) {
	start := time.Now()
	res := cb(ids, ranges)
	end := time.Now()
	diff := end.Sub(start)
	fmt.Printf("%s: %d, duration: %d\n", label, res, diff.Microseconds())
}

func PartA(ids []int, ranges []IdRange) int {
	fresh := NewSet()
	for _, id := range ids {
		for _, idRange := range ranges {
			if idRange.Has(id) {
				fresh.Add(id)

			}
		}
	}
	return fresh.Size()
}

func PartB(ids []int, ranges []IdRange) (freshIds int) {
	// merge ranges when max of next range is inside current range
	merged := mergeSortedRanges(ranges)
	for _, r := range merged {
		freshIds += r.Max - r.Min + 1
	}
	return freshIds
}

func getInputs(lines []string) (ranges []IdRange, ids []int) {
	isIds := false
	for _, line := range lines {
		if line == "" {
			isIds = true
		} else {
			if isIds {
				id, err := strconv.Atoi(line)
				if err != nil {
					log.Fatalf("problem reading id: %v\n", err)
				}
				ids = append(ids, id)
			} else {
				ranges = append(ranges, parseRanges(line))
			}
		}

	}
	return ranges, ids
}

func parseRanges(s string) IdRange {
	sl := strings.Split(s, "-")

	minVal, err := strconv.Atoi(sl[0])
	if err != nil {
		log.Fatalf("problem parsing range minimum: %v", err)
	}

	maxVal, err := strconv.Atoi(sl[1])
	if err != nil {
		log.Fatalf("problem parsing range maximum: %v", err)
	}

	return IdRange{Min: minVal, Max: maxVal}
}

func sortByMin(a IdRange, b IdRange) int {
	return a.Min - b.Min
}

func mergeSortedRanges(ranges []IdRange) []IdRange {
	slices.SortFunc(ranges, sortByMin)
	for {
		count := len(ranges)
		for i := 0; i < len(ranges)-1; i++ {
			for j := i + 1; j < len(ranges); j++ {
				if ranges[i].Overlaps(ranges[j]) {
					ranges[i].Merge(ranges[j])
					ranges = deleteFromListAt(ranges, j)
				}
			}
		}
		if count == len(ranges) {
			break
		}
	}
	return ranges
}

func deleteFromListAt(list []IdRange, idx int) []IdRange {
	return append(list[:idx], list[idx+1:]...)
}
