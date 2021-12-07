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
		currentPosition := allPositions[i]

		for j := 0; j < len(allPositions); j++ {
			if i == j {
				continue
			}

			if allPositions[j] > currentPosition {
				total += (allPositions[j] - currentPosition)

			} else {
				total += (currentPosition - allPositions[j])
			}
		}

		if result == 0 || total < result {
			result = total
		}
	}

	fmt.Println(result)

}
