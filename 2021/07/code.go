package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
		log.Fatalln(err)
	}

	fileScanner := bufio.NewScanner(fh)

	for fileScanner.Scan() {
		inputData = append(inputData, fileScanner.Text())
	}

	return
}

func getCrabs(inputData string) (crabs map[int]int) {

	crabs = make(map[int]int)
	tmp := strings.Split(inputData, ",")

	for i := 0; i < len(tmp); i++ {
		crab, _ := strconv.Atoi(tmp[i])

		_, keyExists := crabs[crab]
		if keyExists {
			crabs[crab] += 1
		} else {
			crabs[crab] = 1
		}
	}

	return
}

func calcFuelUsage(crabs map[int]int, constantFuel bool) (fuel map[int]int) {
	fuel = make(map[int]int)

	min := math.MaxInt
	max := math.MinInt

	for pos, _ := range crabs {
		if pos < min {
			min = pos
		}
		if pos > max {
			max = pos
		}
	}

	for targetPos := min; targetPos <= max; targetPos++ {
		for crabPos, crabAmount := range crabs {

			distance := crabPos - targetPos
			if distance < 0 {
				distance *= -1
			}

			var tmp int

			if !constantFuel {
				tmp = distance * (distance + 1) / 2
			} else {
				tmp = distance
			}

			fuel[targetPos] += (tmp * crabAmount)
		}
	}

	return
}

func getMinFuelUsage(fuel map[int]int) (result int) {
	result = math.MaxInt

	for _, usage := range fuel {
		if usage < result {
			result = usage
		}
	}

	return
}

func part1(crabs map[int]int) (result int) {
	fuel := calcFuelUsage(crabs, true)
	result = getMinFuelUsage(fuel)

	return
}

func part2(crabs map[int]int) (result int) {
	fuel := calcFuelUsage(crabs, false)
	result = getMinFuelUsage(fuel)

	return
}

func main() {
	inputFile := parseArgs()
	inputData := getData(inputFile)

	crabs := getCrabs(inputData[0])

	result1 := part1(crabs)
	fmt.Printf("Part 1: minimum fuel usage: %d\n", result1)

	result2 := part2(crabs)
	fmt.Printf("Part 2: minimum fuel usage: %d\n", result2)

}
