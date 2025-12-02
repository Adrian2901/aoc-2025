package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

// Read the input data file
func readInput() []Range {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Could not read input.txt")
		return nil
	}
	defer file.Close()

	var ranges []Range
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		lineRanges := strings.SplitSeq(line, ",")

		for r := range lineRanges {
			values := strings.Split(r, "-")
			start, err := strconv.Atoi(values[0])
			if err != nil {
				continue
			}

			end, err := strconv.Atoi(values[1])
			if err != nil {
				continue
			}

			ranges = append(ranges, Range{
				Start: start,
				End:   end,
			})
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil
	}

	return ranges
}

func main() {
	input := readInput()
	var invalidNumbers []string
	for _, productRange := range input {
		for i := productRange.Start; i <= productRange.End; i++ {
			idString := strconv.Itoa(i)
			length := len(idString)
			for chunkSize := 1; chunkSize <= length/2; chunkSize++ {
				if length%chunkSize != 0 {
					continue
				}
				chunk := idString[0:chunkSize]
				isInvalid := true
				for i := range length / len(chunk) {
					startIndex := i * chunkSize
					if idString[startIndex:startIndex+chunkSize] != chunk {
						isInvalid = false
					}
				}
				if isInvalid {
					invalidNumbers = append(invalidNumbers, idString)
					break
				}
			}
		}
	}
	sum := 0
	for _, id := range invalidNumbers {
		i, _ := strconv.Atoi(id)
		sum += i
	}
	fmt.Println(sum)
}
