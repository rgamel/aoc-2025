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

	sum := 0
	for _, r := range ranges {
		start, end := ParseRange(r)

		i := start
		for i <= end {
			sum = AddInvalidId(i, sum)
			i++
		}
	}

	fmt.Printf("day 1 sum: %d\n", sum)
}

func IsInvalidId(id int) bool {
	s := strconv.Itoa(id)
	mid := len(s) / 2
	return s[mid:] == s[:mid]
}

func AddInvalidId(id int, sum int) int {
	if IsInvalidId(id) {
		return sum + id
	}
	return sum
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
