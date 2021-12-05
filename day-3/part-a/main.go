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

	commonValues := make(map[int]int)

	for scanner.Scan() {

		var input = scanner.Text()

		for i, c := range input {
			_, exists := commonValues[i]

			if !exists {
				commonValues[i] = 0
			}

			if c == '0' {
				commonValues[i] = commonValues[i] - 1
			} else {
				commonValues[i] = commonValues[i] + 1
			}
		}

	}

	var gammaValues []rune
	var epsilonValues []rune

	for i := 0; i < len(commonValues); i++ {

		value := commonValues[i]

		if value > 0 {
			gammaValues = append(gammaValues, '1')
			epsilonValues = append(epsilonValues, '0')
		} else {
			gammaValues = append(gammaValues, '0')
			epsilonValues = append(epsilonValues, '1')
		}
	}

	gammaResult, _ := strconv.ParseInt(string(gammaValues), 2, 32)
	epsilonResult, _ := strconv.ParseInt(string(epsilonValues), 2, 32)

	result := gammaResult * epsilonResult

	fmt.Printf("Result: %d", result)
}
