package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	count := 0

	for scanner.Scan() {
		outputValues := strings.Split(strings.Split(scanner.Text(), "|")[1], " ")

		for _, outputValue := range outputValues {
			if outputValue == "" {
				continue
			}

			validLengths := []int{2, 3, 4, 7}

			for _, length := range validLengths {
				if len(outputValue) == length {
					count++
				}
			}
		}
	}

	fmt.Println(count)

}
