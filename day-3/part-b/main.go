package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func loadInput() [][]int {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	results := make([][]int, 0)

	for scanner.Scan() {

		var input = scanner.Text()
		posResult := make([]int, 0)

		for _, c := range input {
			value, _ := strconv.Atoi(string(c))
			posResult = append(posResult, value)
		}

		results = append(results, posResult)
	}

	return results
}

func main() {
	input := loadInput()

	oxygenRating := arrayToValue(getRating(input, true))
	c02Rating := arrayToValue(getRating(input, false))

	fmt.Println(c02Rating * oxygenRating)
}

func getRating(input [][]int, mostCommon bool) []int {

	searchValues := input
	searchIndex := 0

	for len(searchValues) > 1 {

		secondaryValues := make([][]int, 0)
		primaryValues := make([][]int, 0)

		for _, value := range searchValues {
			if value[searchIndex] == 0 {
				secondaryValues = append(secondaryValues, value)
			} else {
				primaryValues = append(primaryValues, value)
			}
		}

		if len(primaryValues) >= len(secondaryValues) {
			if mostCommon {
				searchValues = primaryValues
			} else {
				searchValues = secondaryValues
			}
		} else {
			if mostCommon {
				searchValues = secondaryValues

			} else {
				searchValues = primaryValues
			}
		}
		searchIndex += 1
	}

	return searchValues[0]

}

func arrayToValue(input []int) int64 {

	var output string

	for _, i := range input {
		output += strconv.Itoa(i)
	}

	result, _ := strconv.ParseInt(output, 2, 32)

	return result
}
