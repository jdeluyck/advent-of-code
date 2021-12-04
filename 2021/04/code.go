package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseArgs() (inputFile string) {
	if len(os.Args) == 2 {
		inputFile = os.Args[1]
	} else {
		log.Fatalf("Usage: %v INPUTFILE", os.Args[0])
	}

	return
}

func getData(inputFile string) (inputData []string) {
	fh, err := os.Open(inputFile)

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(fh)

	for fileScanner.Scan() {
		inputData = append(inputData, fileScanner.Text())
	}

	return
}

func getNumbers(inputData string) (numbers []int) {

	for _, number := range strings.Split(inputData, ",") {
		tmp, _ := strconv.Atoi(number)
		numbers = append(numbers, tmp)
	}

	return
}

func getBoards(inputData []string) (boards [][][]int) {

	var lineIdx int
	boardCount := -1

	for _, inputLine := range inputData {
		if len(inputLine) == 0 {
			// new board
			boardCount += 1
			var board [][]int
			boards = append(boards, board)
			lineIdx = -1
		} else {
			// line with numbers
			for colIdx, col := range strings.Fields(inputLine) {
				if colIdx == 0 {
					// new line of this board
					var line []int
					boards[boardCount] = append(boards[boardCount], line)
					lineIdx += 1
				}
				tmp, _ := strconv.Atoi(col)
				boards[boardCount][lineIdx] = append(boards[boardCount][lineIdx], tmp)
			}
		}
	}

	return
}

func getMarkBoards(boards [][][]int) (markBoards [][][]string) {

	for boardIdx, board := range boards {
		var markBoard [][]string
		markBoards = append(markBoards, markBoard)

		for lineIdx, line := range board {
			var boardLine []string
			markBoards[boardIdx] = append(markBoards[boardIdx], boardLine)
			markBoards[boardIdx][lineIdx] = make([]string, len(line))
		}
	}

	return
}

func main() {
	inputFile := parseArgs()
	inputData := getData(inputFile)

	numbers := getNumbers(inputData[0])
	boards := getBoards(inputData[1:])
	markBoards := getMarkBoards(boards)

	result1 := part1(numbers, boards, markBoards)
	result2 := part2(numbers, boards, markBoards)

	fmt.Printf("Part 1: winning number: %v, sum: %v, result: %v\n", result1[0], result1[1], result1[2])
	fmt.Printf("Part 2: winning number: %v, sum: %v, result: %v\n", result2[0], result2[1], result2[2])
}

func checkLine(markBoard [][]string, line int, col int) (bingo bool) {
	bingo = true

	// check line
	for y := 0; y < len(markBoard[line]); y++ {
		if markBoard[line][y] != "x" {
			// no bingo possible, exit this loop
			bingo = false
			break
		}
	}

	if bingo == false {
		// check col
		bingo = true
		for x := 0; x < len(markBoard); x++ {
			if markBoard[x][col] != "x" {
				// no bingo possible
				bingo = false
				break
			}
		}
	}

	return
}

func sumBoard(board [][]int, markBoard [][]string) (result int) {
	for x, line := range markBoard {
		for y, val := range line {
			if val != "x" {
				result += board[x][y]
			}
		}
	}

	return
}

func part1(numbers []int, boards [][][]int, markBoards [][][]string) (results [3]int) {

boardcheck:
	for _, number := range numbers {
		for boardIdx, board := range boards {
			for lineIdx, line := range board {
				for colIdx, val := range line {
					if val == number {
						markBoards[boardIdx][lineIdx][colIdx] = "x"

						if checkLine(markBoards[boardIdx], lineIdx, colIdx) == true {
							boardSum := sumBoard(boards[boardIdx], markBoards[boardIdx])
							results = [3]int{number, boardSum, boardSum * number}

							break boardcheck
						}
					}
				}
			}
		}
	}

	return
}

func part2(numbers []int, boards [][][]int, markBoards [][][]string) (results [3]int) {

	winningBoards := make([]int, len(boards))
	winningNumbers := make([]int, len(boards))

	for idx := 0; idx < len(boards); idx++ {
		winningBoards[idx] = -1
	}

	boardWinCount := 0
	for _, number := range numbers {
		for boardIdx, board := range boards {
			if winningBoards[boardIdx] != -1 {
				continue
			}
			for lineIdx, line := range board {
				for colIdx, val := range line {
					if val == number {
						markBoards[boardIdx][lineIdx][colIdx] = "x"

						if checkLine(markBoards[boardIdx], lineIdx, colIdx) == true {
							if winningBoards[boardIdx] == -1 {
								winningBoards[boardIdx] = boardWinCount
								winningNumbers[boardIdx] = number
								boardWinCount += 1
							}
						}
					}
				}
			}
		}
	}

	winningBoard := -1

	for idx := 0; idx < len(winningBoards); idx++ {
		if winningBoards[idx] > winningBoard {
			winningBoard = idx
		}
	}

	boardSum := sumBoard(boards[winningBoard], markBoards[winningBoard])
	results = [3]int{winningNumbers[winningBoard], boardSum, boardSum * winningNumbers[winningBoard]}

	return
}
