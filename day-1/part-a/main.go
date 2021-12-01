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

	var prevMeasurement int
	measurementIncreaseCount := 0

	for scanner.Scan() {
		var depthMeasurement, _ = strconv.Atoi(scanner.Text())

		if prevMeasurement != 0 && depthMeasurement > prevMeasurement {
			measurementIncreaseCount = measurementIncreaseCount + 1
		}

		prevMeasurement = depthMeasurement
	}

	fmt.Printf("Total depth increases: %d", measurementIncreaseCount)

}
