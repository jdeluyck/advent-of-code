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
	fmt.Printf("Part 1: horizontal: %v, depth: %v, result: %v\n", result1[0], result1[1], result1[2])
	fmt.Printf("Part 2: horizontal: %v, depth: %v, result: %v\n", result2[0], result2[1], result2[2])
}

func part1(data []string) (results [3]int) {
	depth := 0
	horizontal := 0

	for _, tmp := range data {

		tmp2 := strings.Fields(tmp)
		motion := tmp2[0]
		amount, _ := strconv.Atoi(tmp2[1])

		if motion == "forward" {
			horizontal += amount
		} else if motion == "down" {
			depth += amount
		} else if motion == "up" {
			depth -= amount
		}
	}

	results = [3]int{horizontal, depth, horizontal * depth}

	return
}

func part2(data []string) (results [3]int) {
	aim := 0
	horizontal := 0
	depth := 0

	for _, tmp := range data {

		tmp2 := strings.Fields(tmp)
		motion := tmp2[0]
		amount, _ := strconv.Atoi(tmp2[1])

		if motion == "forward" {
			horizontal += amount
			depth += aim * amount
		} else if motion == "down" {
			aim += amount
		} else if motion == "up" {
			aim -= amount
		}
	}

	results = [3]int{horizontal, depth, horizontal * depth}

	return
}
