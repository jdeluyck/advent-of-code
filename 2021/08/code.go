package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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
		log.Fatalln(err)
	}

	fileScanner := bufio.NewScanner(fh)

	for fileScanner.Scan() {
		inputData = append(inputData, fileScanner.Text())
	}

	return
}

func sortStuff(strslice []string) []string {
	for idx, str := range strslice {
		strslice[idx] = sortString(str)
	}
	return strslice
}

func sortString(a string) string {
	tmp := strings.Split(a, "")
	sort.Strings(tmp)
	return strings.Join(tmp, "")
}

func parseData(lines []string) (patterns [][]string, messages [][]string) {
	for _, line := range lines {
		chunks := strings.Split(line, "|")
		patterns = append(patterns, sortStuff(strings.Fields(chunks[0])))
		messages = append(messages, sortStuff(strings.Fields(chunks[1])))
	}

	return
}
func solveEasyNumbers(patterns []string) (remainingPatterns []string, digits [10]string) {
	for _, pattern := range patterns {
		switch len(pattern) {
		case 2:
			digits[1] = pattern
		case 3:
			digits[7] = pattern
		case 4:
			digits[4] = pattern
		case 7:
			digits[8] = pattern
		default:
			remainingPatterns = append(remainingPatterns, pattern)
		}
	}

	return
}

func containsPattern(haystack string, needle string) (result bool) {

	if len(needle) == 0 {
		return false
	}

	for _, char := range needle {
		if !strings.ContainsRune(haystack, char) {
			return false
		}
	}

	return true
}

func solveHardNumbers(digits [10]string, patterns []string) ([]string, [10]string) {

	var newPatterns []string

	if len(patterns) == 1 {
		digits[2] = patterns[0]
	} else {

		for _, pattern := range patterns {
			switch len(pattern) {
			case 5:
				if containsPattern(pattern, digits[1]) {
					// len = 5, contains 1 == 3
					digits[3] = pattern
				} else if containsPattern(digits[6], pattern) {
					// len = 5, contains 6 == 5
					digits[5] = pattern
				} else {
					newPatterns = append(newPatterns, pattern)
				}

			case 6:
				if !containsPattern(pattern, digits[1]) {
					// len 6, does not contain 1 == 6
					digits[6] = pattern
				} else if containsPattern(pattern, digits[4]) {
					// len 6, contains 4 == 9
					digits[9] = pattern
				} else {
					digits[0] = pattern
				}
			}
		}
	}
	return newPatterns, digits

}

func main() {
	inputFile := parseArgs()
	inputData := getData(inputFile)

	var allDigits [][10]string

	allPatterns, allMessages := parseData(inputData)

	for _, pattern := range allPatterns {
		patterns, digits := solveEasyNumbers(pattern)

		for len(patterns) > 0 {
			patterns, digits = solveHardNumbers(digits, patterns)
		}

		allDigits = append(allDigits, digits)
	}

	result1 := part1(allMessages, allDigits)
	fmt.Printf("Part 1: simple parsing sum: %d\n", result1)

	result2 := part2(allMessages, allDigits)
	fmt.Printf("Part 2: all parsing sum: %d\n", result2)
}

func part1(messages [][]string, digits [][10]string) (result int) {
	for idx := 0; idx < len(messages); idx++ {
		for _, val := range []string{digits[idx][1], digits[idx][4], digits[idx][7], digits[idx][8]} {
			for _, digit := range messages[idx] {
				if digit == val {
					result += 1
				}
			}
		}
	}
	return
}

func part2(messages [][]string, digits [][10]string) (result int) {
	for idx := 0; idx < len(messages); idx++ {
		for idx2 := 0; idx2 < 4; idx2++ {
			for digitIdx, val := range digits[idx] {
				digit := messages[idx][idx2]
				if digit == val {
					result += digitIdx * int(math.Pow(10, float64(3-idx2)))
				}
			}

		}
	}
	return
}
