package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func main() {
	lines, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	var ranges []*Range

	for _, line := range strings.Split(string(lines), "\n") {
		line = strings.Replace(line, "\r", "", -1)
		if line == "" {
			break
		}
		ranges = append(ranges, parseRange(line))
	}

	// Sort ranges by start value
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	// Merge overlapping ranges
	merged := []*Range{}
	for _, r := range ranges {
		if len(merged) == 0 || merged[len(merged)-1].End < r.Start-1 {
			// No overlap, add as new range
			merged = append(merged, &Range{Start: r.Start, End: r.End})
		} else {
			// Overlapping or adjacent, extend the current range
			if r.End > merged[len(merged)-1].End {
				merged[len(merged)-1].End = r.End
			}
		}
	}

	// Count total valid numbers from merged ranges
	validNumbers := 0
	for _, r := range merged {
		validNumbers += r.End - r.Start + 1
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
