package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// Something about ranges?
// Input: 11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
// 1698522-1698528,446443-446449,38593856-38593862,565653-565659,
// 824824821-824824827,2121212118-2121212124
// THe input will be 1 long list seperated by commas
// Each range is first ID and last ID
// Invalid IDs:
// * No repeating numbers
// * No leading zeroes
// Find all invalid IDs

// 11-22 has two invalid IDs, 11 and 22.
// 95-115 has one invalid ID, 99.
// 998-1012 has one invalid ID, 1010.
// 1188511880-1188511890 has one invalid ID, 1188511885.
// 222220-222224 has one invalid ID, 222222.
// 1698522-1698528 contains no invalid IDs.
// 446443-446449 has one invalid ID, 446446.
// 38593856-38593862 has one invalid ID, 38593859.
// The rest of the ranges contain no invalid IDs.

// Add them all up

func main() {
	invalidIDSum := 0
	lines, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	for _, line := range strings.Split(string(lines), "\n") {
		if line == "" {
			continue
		}
		line = strings.Replace(line, "\r", "", -1)
		ranges := strings.Split(line, ",")
		for _, r := range ranges {
			split := strings.Split(r, "-")
			start, err := strconv.Atoi(split[0])
			if err != nil {
				log.Fatalf("Error converting start to int: %v", err)
			}
			end, err := strconv.Atoi(split[1])
			if err != nil {
				log.Fatalf("Error converting end to int: %v", err)
			}

			for i := start; i <= end; i++ {
				if isInvalid(i) {
					log.Printf("Invalid ID: %d", i)
					invalidIDSum += i
				}
			}

			log.Printf("Start: %d, End: %d, Invalid ID Sum: %d", start, end, invalidIDSum)
		}
	}
}

func isInvalid(id int) bool {
	idString := strconv.Itoa(id)
	// if len(idString)%2 != 0 {
	// 	return false
	// }
	// firstHalf := idString[:len(idString)/2]
	// secondHalf := idString[len(idString)/2:]
	// return firstHalf == secondHalf

	// fmt.Printf("idString: %s, len: %d, divide by 2: %d\n", idString, len(idString), len(idString)/2)
	for i := 1; i <= len(idString)/2; i++ {
		firstPart := idString[:i]
		requiredCount := len(idString) / i
		if len(idString)%i != 0 {
			continue
		}
		if strings.Count(idString, firstPart) == requiredCount {
			return true
		}
		// fmt.Printf("idString: %s, First part: %s, count: %d, requiredCount: %d\n", idString, firstPart, strings.Count(idString, firstPart), requiredCount)

	}
	return false
}
