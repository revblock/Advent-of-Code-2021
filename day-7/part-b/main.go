package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	data, _ := ioutil.ReadFile("input.txt")

	var allPositions []int

	for _, v := range strings.Split(string(data), ",") {
		position, _ := strconv.Atoi(v)
		allPositions = append(allPositions, position)
	}

	result := 0

	for i := 0; i < len(allPositions); i++ {

		total := 0

		for j := 0; j < len(allPositions); j++ {
			difference := 0

			if allPositions[j] > i {
				difference = (allPositions[j] - i)

			} else {
				difference = (i - allPositions[j])
			}

			fuelUsage := 0
			for k := 0; k < difference; k++ {
				fuelUsage += k + 1
			}

			total += fuelUsage
		}

		if result == 0 || total < result {
			result = total
		}
	}

	fmt.Println(result)

}
