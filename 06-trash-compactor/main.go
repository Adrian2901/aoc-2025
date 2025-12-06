package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() ([][]int, []string) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Could not read input.txt")
		return nil, nil
	}
	defer file.Close()

	var numbers [][]int
	var symbols []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		fields := strings.Fields(line)

		if strings.ContainsAny(fields[0], "+*") {
			symbols = append(symbols, fields...)
			return numbers, symbols
		}

		for i, v := range fields {
			num, _ := strconv.Atoi(v)
			for len(numbers) <= i {
				numbers = append(numbers, []int{})
			}
			numbers[i] = append(numbers[i], num)
		}
	}
	return numbers, symbols
}

func main() {
	numbers, symbols := readInput()

	sum := 0

	for i, col := range numbers {
		operator := symbols[i]
		problemResult := 0

		if operator == "+" {
			for _, n := range col {
				problemResult += n
			}
		} else if operator == "*" {
			problemResult = 1
			for _, n := range col {
				problemResult *= n
			}
		}

		sum += problemResult
	}

	fmt.Println(sum)
}
