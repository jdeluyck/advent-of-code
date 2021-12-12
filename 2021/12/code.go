package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

type caveMap map[string][]string

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

func parseData(inputData []string) (outputData caveMap) {

	outputData = make(caveMap)

	for _, line := range inputData {
		tmp := strings.Split(line, "-")
		addPath(outputData, tmp[0], tmp[1])
	}
	return
}

func addPath(caves caveMap, srcCave string, destCave string) caveMap {
	caves[srcCave] = append(caves[srcCave], destCave)
	caves[destCave] = append(caves[destCave], srcCave)

	return caves
}

func searchPaths(caves caveMap, currentMap []string, visited map[string]int, maxVisits int) [][]string {
	pos := currentMap[len(currentMap)-1]

	if pos == "end" {
		return [][]string{currentMap}
	}

	visited[pos]++

	paths := [][]string{}

	for _, nextPos := range caves[pos] {
		if !caveVisited(nextPos, visited, maxVisits) {
			tmp := copyMap(visited)
			nextCave := searchPaths(caves, append(currentMap, nextPos), tmp, maxVisits)
			paths = append(paths, nextCave...)
		}
	}

	return paths
}

func smallCave(caveName string) bool {
	return !unicode.IsUpper([]rune(caveName)[0])
}

func copyMap(visitedMap map[string]int) (copiedMap map[string]int) {
	copiedMap = make(map[string]int)

	for cave, visits := range visitedMap {
		copiedMap[cave] = visits
	}

	return

}

func caveVisited(caveName string, visited map[string]int, maxVisits int) bool {

	if !smallCave(caveName) || visited[caveName] == 0 {
		// big cave or never visited cave, come again
		return false
	}

	if caveName == "start" || caveName == "end" || maxVisits < 2 {
		// not going back to start, end.
		return true
	}

	for idx, caveVisitCount := range visited {
		if smallCave(idx) && caveVisitCount >= maxVisits {
			// small cave and we visited it too many times
			return true
		}
	}

	return false
}

func startSearch(caves caveMap, maxVisits int) (result int) {
	tmp := make(map[string]int)
	result = len(searchPaths(caves, []string{"start"}, tmp, maxVisits))

	return
}

func main() {
	inputFile := parseArgs()
	inputData := parseData(getData(inputFile))

	result1 := startSearch(inputData, 1)
	fmt.Printf("Part 1: Total paths: %d\n", result1)

	result2 := startSearch(inputData, 2)
	fmt.Printf("Part 2: Total paths: %d\n", result2)
}
