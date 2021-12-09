package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	heightMap := make([][]int, 0)

	for scanner.Scan() {
		input := scanner.Text()

		line := make([]int, 0)

		for _, value := range input {
			parsedValue, _ := strconv.Atoi(string(value))
			line = append(line, parsedValue)
		}

		heightMap = append(heightMap, line)
	}

	result := 0

	for rowIndex, row := range heightMap {

		for colIndex, value := range row {

			up := 9
			right := 9
			down := 9
			left := 9

			// Up value
			if rowIndex != 0 {
				up = heightMap[rowIndex-1][colIndex]
			}

			// Left value
			if colIndex != 0 {
				left = row[colIndex-1]
			}

			// Right value
			if colIndex != len(row)-1 {
				right = row[colIndex+1]
			}

			// Down value
			if rowIndex != len(heightMap)-1 {
				down = heightMap[rowIndex+1][colIndex]
			}

			if value < up && value < right && value < down && value < left {
				result = result + 1 + value
			}

		}

	}

	fmt.Println(result)

}
