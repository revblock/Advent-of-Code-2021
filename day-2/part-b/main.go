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
	var aim int

	for scanner.Scan() {
		var input = strings.Split(scanner.Text(), " ")
		var command = input[0]
		var value, _ = strconv.Atoi(input[1])

		if command == "forward" {
			horizontal = horizontal + value
			depth = depth + aim*value
		}

		if command == "down" {
			aim = aim + value
		}

		if command == "up" {
			aim = aim - value
		}

	}

	fmt.Printf("Result: %d", horizontal*depth)

}
