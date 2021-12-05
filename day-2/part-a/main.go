package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var horizontal int
	var depth int

	for scanner.Scan() {
		var input = strings.Split(scanner.Text(), " ")
		var command = input[0]
		var value, _ = strconv.Atoi(input[1])

		if command == "forward" {
			horizontal = horizontal + value
		}

		if command == "down" {
			depth = depth + value
		}

		if command == "up" {
			depth = depth - value
		}

	}

	fmt.Printf("Result: %d", horizontal*depth)

}
