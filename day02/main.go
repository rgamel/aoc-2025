package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("error opening file")
	}
	defer f.Close()

	ranges := []string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		r := strings.Split(line, ",")
		for _, v := range r {
			ranges = append(ranges, v)
		}

	}

	sumA := 0
	sumB := 0
	for _, r := range ranges {
		start, end := ParseRange(r)

		i := start
		for i <= end {
			sumA = AddInvalidId(i, sumA, IsInvalidId)
			sumB = AddInvalidId(i, sumB, IsInvalidIdB)
			i++
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

func ParseRange(r string) (start int, end int) {
	v := strings.Split(r, "-")
	start, err := strconv.Atoi(v[0])
	if err != nil {
		log.Fatal("problem parsing ranges")
	}

	end, err = strconv.Atoi(v[1])
	if err != nil {
		log.Fatal("problem parsing ranges")
	}

	return start, end
}

func IsInvalidIdB(id int) bool {
	sId := strconv.Itoa(id)
	return strings.Index((sId + sId)[1:], sId) != len(sId)-1
}
