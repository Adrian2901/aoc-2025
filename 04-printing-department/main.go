package main

import (
	"bufio"
	"fmt"
	"os"
)

// Read the input data file
func readInput() [][]string {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Could not read input.txt")
		return nil
	}
	defer file.Close()

	var grid [][]string
	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		row := make([]string, len(line))
		for i, char := range line {
			row[i] = string(char)
		}
		grid = append(grid, row)
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil
	}

	return grid
}

func main() {
	input := readInput()
	var count int
	for {
		countIteration := 0
		for i, row := range input {
			for j, v := range row {
				if v == "@" {
					if countRolls(input, i, j) < 4 {
						countIteration++
						input[i][j] = "."
					}
				}
			}
		}
		if countIteration == 0 {
			break
		} else {
			count += countIteration
		}
	}

	fmt.Println(count)
}

func countRolls(grid [][]string, row int, col int) int {
	count := 0
	rows := len(grid)
	cols := len(grid[0])
	if row > 0 {
		if col > 0 && col < cols-1 {
			for i := -1; i <= 1; i++ {
				if grid[row-1][col+i] == "@" {
					count++
				}
			}
		} else if col > 0 {
			for i := -1; i <= 0; i++ {
				if grid[row-1][col+i] == "@" {
					count++
				}
			}
		} else {
			for i := 0; i <= 1; i++ {
				if grid[row-1][col+i] == "@" {
					count++
				}
			}
		}
	}
	if row < rows-1 {
		if col > 0 && col < cols-1 {
			for i := -1; i <= 1; i++ {
				if grid[row+1][col+i] == "@" {
					count++
				}
			}
		} else if col > 0 {
			for i := -1; i <= 0; i++ {
				if grid[row+1][col+i] == "@" {
					count++
				}
			}
		} else {
			for i := 0; i <= 1; i++ {
				if grid[row+1][col+i] == "@" {
					count++
				}
			}
		}
	}
	if col > 0 {
		if grid[row][col-1] == "@" {
			count++
		}
	}
	if col < cols-1 {
		if grid[row][col+1] == "@" {
			count++
		}
	}
	return count
}
