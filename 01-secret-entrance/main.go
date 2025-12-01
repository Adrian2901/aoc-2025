package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	StartPosition                = 50
	LowerBoundary, UpperBoundary = 0, 99
)

type Instruction struct {
	Direction string
	Count     int
}

// Read the input data file
func readInput() []Instruction {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Could not read input.txt")
		return nil
	}
	defer file.Close()

	var instructions []Instruction
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		// Extract direction (first character) and count (remaining characters)
		direction := string(line[0])
		countStr := line[1:]

		count, err := strconv.Atoi(countStr)
		if err != nil {
			fmt.Printf("Error parsing count from line '%s': %v\n", line, err)
			continue
		}

		instructions = append(instructions, Instruction{
			Direction: direction,
			Count:     count,
		})
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil
	}

	return instructions
}

func main() {
	instructions := readInput()
	position := StartPosition
	zeroesCount := 0
	for _, rotation := range instructions {
		switch rotation.Direction {
		case "L":
			position -= rotation.Count
		case "R":
			position += rotation.Count
		default:
			continue
		}
		if position < LowerBoundary {
			position = (UpperBoundary + 1) + (position % 100)
		}
		if position > UpperBoundary {
			position = position % 100
		}
		if position == 0 {
			zeroesCount++
		}
	}
	fmt.Printf("Password: %d ", zeroesCount)
}
