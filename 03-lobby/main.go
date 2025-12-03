package main

import (
	"bufio"
	"fmt"
	"math"
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

func toInt(character byte) int {
	// to cast an ASCII byte to number we deduct 48
	return int(character) - 48
}

func main() {
	banks := readInput()
	var numbers [12]int
	sum := 0
	for _, bank := range banks {
		bankSize := len(bank)
		initialNumbers := bank[bankSize-12:]
		for i := range 12 {
			numbers[i] = toInt(initialNumbers[i])
		}
		for i := bankSize - 13; i >= 0; i-- {
			num := toInt(bank[i])
			if num < numbers[0] {
				continue
			} else {
				temp := numbers[0]
				numbers[0] = num
				for i := 1; i < 12; i++ {
					if temp >= numbers[i] {
						newValue := temp
						temp = numbers[i]
						numbers[i] = newValue
					} else {
						break
					}
				}
			}
		}
		for i := range 12 {
			sum += numbers[i] * int(math.Pow(float64(10), float64(11-i)))
		}
	}
	fmt.Println(sum)
}
