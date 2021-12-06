package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	x int
	y int
}

type bearing struct {
	path       []coordinate
	isDiagonal bool
}

func newBearing(x1 int, y1 int, x2 int, y2 int) bearing {

	var isDiagonal bool
	if x1 == x2 || y1 == y2 {
		isDiagonal = false
	} else {
		isDiagonal = true
	}

	var path []coordinate

	x := x1
	y := y1

	for x != x2 || y != y2 {

		path = append(path, coordinate{x, y})

		if x < x2 {
			x++
		} else if x > x2 {
			x--
		}

		if y < y2 {
			y++
		} else if y > y2 {
			y--
		}
	}

	path = append(path, coordinate{x2, y2})

	return bearing{path, isDiagonal}
}

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	allCoordinates := make(map[string]int)

	for scanner.Scan() {

		rawValue := scanner.Text()
		csv := strings.Replace(rawValue, " -> ", ",", 1)
		values := strings.Split(csv, ",")

		y1, _ := strconv.Atoi(values[1])
		x2, _ := strconv.Atoi(values[2])
		x1, _ := strconv.Atoi(values[0])
		y2, _ := strconv.Atoi(values[3])

		bearing := newBearing(x1, y1, x2, y2)

		if bearing.isDiagonal {
			continue
		}

		for _, coordinate := range bearing.path {
			key := fmt.Sprintf("%d,%d", coordinate.x, coordinate.y)
			c := allCoordinates[key]
			allCoordinates[key] = c + 1
		}
	}

	result := 0

	for _, v := range allCoordinates {
		if v >= 2 {
			result++
		}
	}

	fmt.Println(result)
}
