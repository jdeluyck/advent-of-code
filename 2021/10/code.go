package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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
		log.Fatalln(err)
	}

	fileScanner := bufio.NewScanner(fh)

	for fileScanner.Scan() {
		inputData = append(inputData, fileScanner.Text())
	}

	return
}

func push(array []string, char rune) []string {
	return append(array, string(char))
}

func pop(array []string) ([]string, rune) {
	n := len(array) - 1
	char := array[n]

	return array[:n], rune(char[0])
}

func findScore(inputData []string, charVal map[rune]int, complete bool) (result int) {
	var scores []int

	for _, line := range inputData {
		var nextClose []string

		for _, char := range line {
			// fmt.Printf("Expected closing chars: %v\n", nextClose)

			if char == '{' {
				nextClose = push(nextClose, '}')
			} else if char == '<' {
				nextClose = push(nextClose, '>')
			} else if char == '[' {
				nextClose = push(nextClose, ']')
			} else if char == '(' {
				nextClose = push(nextClose, ')')
			} else {
				n := len(nextClose)
				var nextCloseChar rune

				if n != 0 {
					nextClose, nextCloseChar = pop(nextClose)
				}

				if char != nextCloseChar {
					// corrupt line
					if !complete {
						result += charVal[char]
					}
					nextClose = nil
					break
				}
			}
		}

		if complete && len(nextClose) > 0 {
			var char rune
			var tmp int

			for len(nextClose) > 0 {
				nextClose, char = pop(nextClose)
				tmp = tmp*5 + charVal[char]
			}
			scores = append(scores, tmp)
		}
	}

	if complete {
		sort.Ints(scores)
		tmp := int(math.Ceil(float64(len(scores) / 2)))
		result = scores[tmp]

	}

	return
}

func part1(inputData []string) (result int) {
	charVal := make(map[rune]int, 4)
	charVal[')'] = 3
	charVal[']'] = 57
	charVal['}'] = 1197
	charVal['>'] = 25137

	result = findScore(inputData, charVal, false)

	return
}

func part2(inputData []string) (result int) {
	charVal := make(map[rune]int, 4)
	charVal[')'] = 1
	charVal[']'] = 2
	charVal['}'] = 3
	charVal['>'] = 4

	result = findScore(inputData, charVal, true)

	return
}

func main() {
	inputFile := parseArgs()

	inputData := getData(inputFile)

	result1 := part1(inputData)
	fmt.Printf("Part 1: syntax error score: %d\n", result1)

	result2 := part2(inputData)
	fmt.Printf("Part 2: autocomplete score: %d\n", result2)
}
