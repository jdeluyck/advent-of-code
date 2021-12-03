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

func main() {
	inputFile := parseArgs()
	inputData := getData(inputFile)

	result1 := part1(inputData)
	result2 := part2(inputData)
	fmt.Printf("Part 1: gamma: %v, epsilon: %v, result: %v\n", result1[0], result1[1], result1[2])
	fmt.Printf("Part 2: oxygen: %v, CO2: %v, result: %v\n", result2[0], result2[1], result2[2])
}

func countOccur(data []string) [][]int {
	cols := len(data[0])

	countOccur := make([][]int, cols) // []([2]int)

	for idx := 0; idx < cols; idx++ {
		countOccur[idx] = make([]int, 2)
	}

	for _, line := range data {
		for col, bit := range line {
			bitVal, _ := strconv.Atoi(string(bit))
			countOccur[col][bitVal] += 1
		}
	}

	return countOccur
}

func part2(data []string) (results [3]int64) {
	oxygen, _ := strconv.ParseInt(filterData(data, false, "1"), 2, 64)
	co2, _ := strconv.ParseInt(filterData(data, true, "0"), 2, 64)

	results = [3]int64{oxygen, co2, oxygen * co2}
	return

}

func filterData(data []string, min bool, lookupval string) string {
	rows := len(data)
	cols := len(data[0])

	tmp := data
	for col := 0; col < cols; col++ {

		countOccur := countOccur(tmp)
		needle := lookupval

		if min {
			if countOccur[col][1] < countOccur[col][0] {
				needle = "1"
			}
		} else {
			if countOccur[col][0] > countOccur[col][1] {
				needle = "0"
			}

		}
		tmp = lookFor(tmp, col, needle)

		rows = len(tmp)

		if rows == 1 {
			break
		}
	}

	return tmp[0]

}

func lookFor(data []string, col int, val string) (result []string) {
	for _, line := range data {
		tmp := strings.Split(line, "")

		if tmp[col] == val {
			result = append(result, line)
		}
	}

	return
}

func part1(data []string) (results [3]int64) {
	gammaStr := ""
	epsilonStr := ""

	length := len(data[0])

	count0 := make([]int, length)
	count1 := make([]int, length)

	for _, tmp := range data {
		for idx, bit := range tmp {
			if bit == '0' {
				count0[idx] += 1
			} else {
				count1[idx] += 1
			}
		}
	}

	for idx := 0; idx < len(count0); idx++ {
		if count0[idx] > count1[idx] {
			gammaStr += "0"
			epsilonStr += "1"
		} else {
			gammaStr += "1"
			epsilonStr += "0"
		}
	}

	gamma, _ := strconv.ParseInt(gammaStr, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonStr, 2, 64)

	results = [3]int64{gamma, epsilon, gamma * epsilon}

	return

}
