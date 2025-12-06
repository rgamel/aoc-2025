package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadInput(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("problem reading input file: %v\n", err)
	}
	defer f.Close()
	lines := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}
	return lines
}
