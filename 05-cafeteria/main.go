package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

// Read the input data file
func readInput() ([]Range, []int) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Could not read input.txt")
		return nil, nil
	}
	defer file.Close()

	var ranges []Range
	var ingredients []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		ids := strings.Split(line, "-")
		startID, _ := strconv.Atoi(ids[0])
		endID, _ := strconv.Atoi(ids[1])

		ranges = append(ranges, Range{Start: startID, End: endID})
	}

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		id, _ := strconv.Atoi(line)

		ingredients = append(ingredients, id)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil, nil
	}

	return ranges, ingredients
}

func main() {
	ranges, ingredients := readInput()
	count := 0
	slices.SortFunc(ranges, cmp)
	for _, ingredient := range ingredients {
		if isFresh(ranges, ingredient) {
			count++
		}
	}
	fmt.Println(count)
	var freshIds, tempStart, tempEnd int
	for _, idRange := range ranges {
		if idRange.Start == tempStart {
			continue
		} else {
			if idRange.Start <= tempEnd {
				if idRange.End > tempEnd {
					freshIds += idRange.End - tempEnd
					tempEnd = idRange.End
				}
			} else {
				freshIds += idRange.End - idRange.Start + 1
				tempEnd = idRange.End
			}
			tempStart = idRange.Start
			tempEnd = idRange.End
		}
	}
	fmt.Println(freshIds)
}

func cmp(a, b Range) int {
	if a.Start < b.Start {
		return -1
	} else if a.Start > b.Start {
		return 1
	} else {
		if a.End > b.End {
			return -1
		} else if a.End < b.End {
			return 1
		} else {
			return 0
		}
	}
}

func isFresh(ranges []Range, id int) bool {
	left, right := 0, len(ranges)-1
	result := -1

	for left <= right {
		mid := (left + right) / 2
		if ranges[mid].Start <= id {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	for i := result; i >= 0; i-- {
		if ranges[i].Start <= id && ranges[i].End >= id {
			return true
		}
	}

	return false
}
