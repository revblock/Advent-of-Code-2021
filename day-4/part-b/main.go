package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bingoValue struct {
	value  string
	marked bool
}

func generateWinningCombinations(rows int, columns int) [][]int {

	winningRanges := make([][]int, 0)

	for c := 0; c < columns; c++ {

		winningRangeX := make([]int, 0)
		winningRangeY := make([]int, 0)

		for r := 0; r < rows; r++ {
			winningRangeX = append(winningRangeX, r+(c*columns))
			winningRangeY = append(winningRangeY, c+(r*rows))
		}

		winningRanges = append(winningRanges, winningRangeX, winningRangeY)
	}

	return winningRanges
}

func loadInput() ([]string, [][]bingoValue) {
	file, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(file), "\n")

	calledNumbers := strings.Split(lines[0], ",")

	bingoCards := make([][]bingoValue, 0)

	for i := 1; i < len(lines); i++ {

		bingoLine := lines[i]
		if bingoLine == "" {
			bingoCard := make([]bingoValue, 0)
			bingoCards = append(bingoCards, bingoCard)
		}

		bingoCard := &bingoCards[len(bingoCards)-1]

		for _, cardValue := range strings.Split(bingoLine, " ") {

			if cardValue == "" {
				continue
			}

			*bingoCard = append(*bingoCard, bingoValue{value: cardValue, marked: false})
		}
	}

	return calledNumbers, bingoCards
}

func main() {

	rows := 5
	columns := 5
	winningRanges := generateWinningCombinations(rows, columns)
	calledNumbers, bingoCards := loadInput()

	var lastCard []bingoValue
	var lastNumberIndex int
	var lastNumberCalled int

	for cardIndex := 0; cardIndex < len(bingoCards); cardIndex++ {

	out:
		for numberIndex, calledNumber := range calledNumbers {

			for valueIndex := 0; valueIndex < len(bingoCards[cardIndex]); valueIndex++ {

				cardValue := &bingoCards[cardIndex][valueIndex]

				if calledNumber == cardValue.value {
					cardValue.marked = true

					if checkForWinner(bingoCards[cardIndex], winningRanges) {

						lastCalled, _ := strconv.Atoi(calledNumber)

						if numberIndex > lastNumberIndex {
							lastCard = bingoCards[cardIndex]
							lastNumberIndex = numberIndex
							lastNumberCalled = lastCalled
						}

						break out
					}
				}
			}
		}
	}

	fmt.Println(lastCard)
	fmt.Println(lastNumberCalled)

	result := 0
	for _, i := range lastCard {
		if !i.marked {
			value, _ := strconv.Atoi(i.value)
			result += value
		}
	}

	fmt.Println(result * lastNumberCalled)
}

func checkForWinner(bingoCard []bingoValue, winningRanges [][]int) bool {

	for _, wr := range winningRanges {

		winner := true

		for _, wrv := range wr {

			if !bingoCard[wrv].marked {
				winner = false
			}

		}

		if winner {
			return true
		}

	}

	return false
}
