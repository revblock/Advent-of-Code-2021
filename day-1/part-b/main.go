package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func load() (depthMeasurements []int) {

	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		x, _ := strconv.Atoi(scanner.Text())
		depthMeasurements = append(depthMeasurements, x)
	}
	return depthMeasurements
}

func main() {

	depthMeasurements := load()
	var prevMeasurementSum int
	var increaseCount int

	for i, depthMeasurement := range depthMeasurements {

		if i+2 > len(depthMeasurements)-1 {
			break
		}

		value1 := depthMeasurement
		value2 := depthMeasurements[i+1]
		value3 := depthMeasurements[i+2]

		currentMeasurementSum := value1 + value2 + value3

		if prevMeasurementSum > 0 && currentMeasurementSum > prevMeasurementSum {
			increaseCount = increaseCount + 1
		}

		prevMeasurementSum = currentMeasurementSum
	}

	fmt.Printf("Total depth increases: %d", increaseCount)
}
