package main

import (
	"bufio"
	"fmt"
	"os"
)

// Read the input data file
func readInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Could not read input.txt")
		return nil
	}
	defer file.Close()

	var banks []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		banks = append(banks, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil
	}

	return banks
}

func main() {
	banks := readInput()
	sum := 0
	for _, bank := range banks {
		bankSize := len(bank)
		// to cast an ASCII byte to number we deduct 48
		first := int(bank[bankSize-2]) - 48
		second := int(bank[bankSize-1]) - 48
		for i := bankSize - 3; i >= 0; i-- {
			num := int(bank[i]) - 48
			if num < first {
				continue
			} else {
				if first > second {
					second = first
				}
				first = num
			}
		}
		sum += first*10 + second
	}
	fmt.Println(sum)
}
