package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type XYCoords struct {
	X, Y int
}

type EnergyMap map[XYCoords]int

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

func parseData(inputData []string) (outputData EnergyMap) {

	outputData = make(map[XYCoords]int)

	for y, line := range inputData {
		for x, energy := range line {
			outputData[XYCoords{X: x, Y: y}] = int(energy - '0')
		}
	}

	return
}

func getNeighbours(pos XYCoords) (result [8]XYCoords) {

	N := XYCoords{0, 1}
	NE := XYCoords{1, 1}
	E := XYCoords{1, 0}
	SE := XYCoords{1, -1}
	S := XYCoords{0, -1}
	SW := XYCoords{-1, -1}
	W := XYCoords{-1, 0}
	NW := XYCoords{-1, 1}

	AllNeighs := [8]XYCoords{N, NE, E, SE, S, SW, W, NW}

	for idx, neighbour := range AllNeighs {
		result[idx] = XYCoords{
			X: pos.X + neighbour.X,
			Y: pos.Y + neighbour.Y,
		}
	}

	return
}

func increaseEnergy(squids EnergyMap) (EnergyMap, int) {
	flashed := make(map[XYCoords]bool)

	// increase values
	for pos := range squids {
		squids[pos]++
	}

	for pos := range squids {
		if squids[pos] > 9 {
			squids, flashed = flashSquid(squids, flashed, pos)
		}
	}

	return squids, len(flashed)
}

func flashSquid(squids EnergyMap, flashed map[XYCoords]bool, pos XYCoords) (EnergyMap, map[XYCoords]bool) {
	squids[pos] = 0

	flashed[pos] = true

	for _, tmp := range getNeighbours(pos) {
		if _, ok := squids[tmp]; ok {
			// neighbour exists
			if _, ok2 := flashed[tmp]; !ok2 {
				// it already flashed
				squids[tmp]++

				if squids[tmp] > 9 {
					squids, flashed = flashSquid(squids, flashed, tmp)
				}
			}
		}
	}

	return squids, flashed
}

func part1(squids EnergyMap, rounds int) (EnergyMap, int) {
	var flashes, result int

	for x := 1; x <= rounds; x++ {
		squids, flashes = increaseEnergy(squids)
		result += flashes
	}

	return squids, result
}

func part2(squids EnergyMap, startRound int) (result int) {
	flashes := 0
	result = startRound

	for flashes < 100 {
		squids, flashes = increaseEnergy(squids)
		result++
	}

	return
}

func main() {
	inputFile := parseArgs()
	inputData := parseData(getData(inputFile))

	squids, result1 := part1(inputData, 100)
	fmt.Printf("Part 1: Total flashes: %d\n", result1)

	result2 := part2(squids, 100)
	fmt.Printf("Part 2: first all flash step: %d\n", result2)
}
