package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func (r *Range) Contains(input int) bool {
	return r.Start <= input && r.End >= input
}

func main() {
	lines, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	var lookingForRange = true
	var availableRanges []*Range = []*Range{}
	var validNumbers = 0

	for _, line := range strings.Split(string(lines), "\n") {
		line = strings.Replace(line, "\r", "", -1)
		if line == "" {
			lookingForRange = false
			continue
		}

		if lookingForRange {
			availableRanges = append(availableRanges, parseRange(line))
		} else {
			for _, numberRange := range availableRanges {
				number, err := strconv.Atoi(line)
				if err != nil {
					log.Fatalf("Error converting number to int: %v", err)
				}
				if numberRange.Contains(number) {
					validNumbers++
					break
				}
			}
		}
	}
	fmt.Printf("Valid numbers: %d\n", validNumbers)
}

func parseRange(line string) *Range {
	parts := strings.Split(line, "-")
	start, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatalf("Error converting start to int: %v", err)
	}
	end, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalf("Error converting end to int: %v", err)
	}
	return &Range{Start: start, End: end}
}
