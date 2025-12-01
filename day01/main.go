package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Store struct {
	Count int
}

func (s *Store) Increment() {
	s.Count++
}

func main() {
	lines := []string{}
	curr := 50
	store_a := Store{Count: 0}
	store_b := Store{Count: 0}

	// read input file
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("Problem reading file: %v", err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		next := s.Text()
		lines = append(lines, next)
	}

	for _, line := range lines {
		dir, amt := ParseDirection(line)
		rotFn := RotateRight
		if dir == "L" {
			rotFn = RotateLeft
		}
		curr = rotFn(curr, amt, &store_b)
		CheckForZero(curr, &store_a)
	}

	fmt.Printf("password A: %d\n", store_a.Count)
	fmt.Printf("password B: %d\n", store_b.Count)
}

func ParseDirection(input string) (dir string, amt int) {
	dir = string(input[0])

	amt, err := strconv.Atoi((input[1:]))
	if err != nil {
		log.Fatalf("could not read line: %s", input)
	}

	return dir, amt
}

func RotateRight(start int, dist int, store *Store) (end int) {
	end = start + dist
	for end > 99 {
		store.Increment()
		end -= 100
	}
	return end
}

func RotateLeft(start int, dist int, store *Store) (end int) {
	for dist >= 100 {
		dist -= 100
		store.Increment()
	}

	end = start - dist

	if end < 0 {
		end += 100
	}

	CheckForZero(end, store)

	return end
}

func CheckForZero(input int, store *Store) {
	if input == 0 {
		store.Increment()
	}
}
