package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	utils "aoc2025/shared"
)

type IdRange struct {
	Min, Max int
}

func main() {
	lines := utils.ReadInput("./input.txt")
	ranges := []string{}

	sumA := 0
	sumB := 0
	for _, line := range lines {
		r := strings.Split(line, ",")
		for _, v := range r {
			ranges = append(ranges, v)
			idRange := ParseRange(v)
			i := idRange.Min
			for i <= idRange.Max {
				sumA = AddInvalidId(i, sumA, IsInvalidId)
				sumB = AddInvalidId(i, sumB, IsInvalidIdB)
				i++
			}
		}
	}
	fmt.Printf("part A sum: %d\n", sumA)
	fmt.Printf("part B sum: %d\n", sumB)
}

func IsInvalidId(id int) bool {
	s := strconv.Itoa(id)
	mid := len(s) / 2
	return s[mid:] == s[:mid]
}

func AddInvalidId(id int, sum int, findInvalid func(id int) bool) int {
	addend := 0
	if findInvalid(id) {
		addend = id
	}
	return sum + addend
}

func ParseRange(r string) IdRange {
	v := strings.Split(r, "-")
	start, err := strconv.Atoi(v[0])
	if err != nil {
		log.Fatal("problem parsing ranges")
	}

	end, err := strconv.Atoi(v[1])
	if err != nil {
		log.Fatal("problem parsing ranges")
	}

	return IdRange{Min: start, Max: end}
}

func IsInvalidIdB(id int) bool {
	sId := strconv.Itoa(id)
	return strings.Index((sId + sId)[1:], sId) != len(sId)-1
}
