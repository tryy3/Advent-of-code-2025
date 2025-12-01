package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Find the password
// Dial from 0-99
// A click each time it reaches a number
// Puzzle input a sequence of rotations
// One per line which tells you how to open
// A rotation starts with either L or R and the distance
// Dial starts at 50
// Actual password is the number of times the dial is left pointing at 0

// Example:
// L68
// L30
// R48
// L5
// R60
// L55
// L1
// L99
// R14
// L82
// Output: 3

var (
	currentPosition = 50
	amountOfZeroes  = 0
)

func main() {
	lines, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	for _, line := range strings.Split(string(lines), "\n") {
		line = strings.Replace(line, "\r", "", -1)
		direction := string(line[0])
		distance, err := strconv.Atoi(string(line[1:]))
		if err != nil {
			log.Fatalf("Error converting distance to int: %v", err)
		}
		currentPosition = move(direction, distance)

		if currentPosition == 0 {
			amountOfZeroes++
		}
		fmt.Println(currentPosition)
	}
	fmt.Printf("Amount of zeroes: %d\n", amountOfZeroes)
}

func move(direction string, distance int) int {
	newPosition := currentPosition
	if direction == "L" {
		newPosition = (newPosition - distance + 100) % 100
	} else {
		newPosition = (newPosition + distance) % 100
	}
	return newPosition
}
