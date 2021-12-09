package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"sort"
)

type XYCoords struct {
	X, Y int
}

type HeightMap map[XYCoords]int

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

func parseData(inputData []string) (outputData HeightMap) {

	outputData = make(map[XYCoords]int)

	for y, line := range inputData {
		for x, height := range line {
			outputData[XYCoords{X: x, Y: y}] = int(height - '0')
		}
	}

	return
}

func findLowPoints(inputData HeightMap) (results []XYCoords) {
	for aPos := range inputData {
		if isLowPoint(inputData, aPos) {
			results = append(results, aPos)
		}
	}

	return
}

func getNeighbours(pos XYCoords) (result [4]XYCoords) {

	N := XYCoords{0, 1}
	E := XYCoords{1, 0}
	S := XYCoords{0, -1}
	W := XYCoords{-1, 0}
	AllNeighs := [4]XYCoords{N, E, S, W}

	for idx, neighbour := range AllNeighs {
		result[idx] = XYCoords{
			X: pos.X + neighbour.X,
			Y: pos.Y + neighbour.Y,
		}
	}

	return
}

func isLowPoint(inputData HeightMap, pos XYCoords) bool {
	currHeight, ok := inputData[pos]

	if !ok {
		return false
	}

	neighbours := getNeighbours(pos)

	for _, aPos := range neighbours {
		neighbourHeight, ok := inputData[aPos]
		if !ok {
			continue
		}

		if neighbourHeight <= currHeight {
			return false
		}
	}
	return true
}

func findBasinSize(inputData HeightMap, pos XYCoords) (result int) {
	queue := list.New()
	queue.PushBack(pos)

	visited := map[XYCoords]struct{}{
		pos: {},
	}

	for item := queue.Front(); item != nil; item = item.Next() {
		aPos := item.Value.(XYCoords)

		height, ok := inputData[aPos]

		if !ok || height == 9 {
			continue
		}

		result++

		for _, tmp := range getNeighbours(aPos) {

			if _, ok := visited[tmp]; ok {
				continue
			}

			visited[tmp] = struct{}{}
			queue.PushBack(tmp)
		}
	}

	return
}

func part1(inputData HeightMap, coords []XYCoords) (result int) {
	for _, aPos := range coords {
		result += inputData[aPos] + 1
	}
	return
}

func part2(inputData HeightMap, coords []XYCoords) (result int) {
	result = 1
	basinSizes := make([]int, len(coords))

	for idx, aPos := range coords {
		basinSizes[idx] = findBasinSize(inputData, aPos)
	}

	sort.Ints(basinSizes)

	for _, val := range basinSizes[len(basinSizes)-3:] {
		result *= val
	}

	return

}

func main() {
	inputFile := parseArgs()
	inputData := parseData(getData(inputFile))

	lowpoints := findLowPoints(inputData)

	result1 := part1(inputData, lowpoints)
	fmt.Printf("Part 1: sum risklevels: %d\n", result1)

	result2 := part2(inputData, lowpoints)
	fmt.Printf("Part 2: product basins: %d\n", result2)
}
