package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	ranges, ids := getInputs("./input.txt")
	run(ids, ranges, "Part A", PartA)
	run(ids, ranges, "Part B", PartB)
}

type IdRange struct {
	Min int
	Max int
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

func getInputs(filename string) (ranges []IdRange, ids []int) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("problem reading input file: %v\n", err)
	}
	defer f.Close()

	isIds := false

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
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
			if IsInRange(idRange, id) {
				fresh.Add(id)

			}
		}
	}
	return fresh.Size()
}

func sortByMin(a IdRange, b IdRange) int {
	if a.Min < b.Min {
		return -1
	}

	if a.Min > b.Min {
		return 1
	}
	return 0
}

func PartB(ids []int, ranges []IdRange) int {
	// sort ranges by min
	slices.SortFunc(ranges, sortByMin)
	// merge ranges when max of next range is inside current range
	merged := mergeSortedRanges(ranges)

	total := 0
	for _, r := range merged {
		// do max - min + 1
		total += r.Max - r.Min + 1
	}
	return total
}

func IsInRange(idRange IdRange, id int) bool {
	return id >= idRange.Min && id <= idRange.Max
}

func mergeSortedRanges(ranges []IdRange) []IdRange {
	for {
		count := len(ranges)
		for i := 0; i < len(ranges)-1; i++ {
			for j := i + 1; j < len(ranges); j++ {
				if IsInRange(ranges[i], ranges[j].Min) && IsInRange(ranges[i], ranges[j].Max) {
					// j is completely within i
					// delete it
					ranges = deleteFromListAt(ranges, j)
				} else if IsInRange(ranges[i], ranges[j].Min) && !IsInRange(ranges[i], ranges[j].Max) {
					// overlap with j starting within i and terminating higher than i
					// max i's max equal j's, then delete it
					ranges[i].Max = ranges[j].Max
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
