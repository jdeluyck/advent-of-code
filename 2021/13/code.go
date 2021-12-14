package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type XYCoords struct {
	X, Y int
}

type fold struct {
	Orientation rune
	Position    int
}

type pointMap map[XYCoords]bool

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

func parseData(inputData []string) (outputData pointMap, folds []fold) {

	outputData = make(pointMap)

	var y, x, pos int
	var dir rune

	for _, line := range inputData {
		if len(line) == 0 {
			continue
		} else if strings.HasPrefix(line, "fold") {
			cnt, err := fmt.Sscanf(line, "fold along %c=%d", &dir, &pos)

			if err != nil {
				log.Fatalln(err)
			}

			if cnt != 2 {
				log.Fatalln("Unexpected string")
			}

			folds = append(folds, fold{Orientation: dir, Position: pos})

		} else {
			cnt, err := fmt.Sscanf(line, "%d,%d", &x, &y)

			if err != nil {
				log.Fatalln(err)
			}

			if cnt != 2 {
				log.Fatalln("Unexpected string!")
			}

			outputData[XYCoords{X: x, Y: y}] = true

		}
	}

	return outputData, folds
}

func doFold(points pointMap, fold fold) pointMap {
	var tmp []XYCoords

	if fold.Orientation == 'x' {
		for point := range points {
			if point.X > fold.Position {
				tmp = append(tmp, point)
			}
		}

		for _, point := range tmp {
			delete(points, point)
			point.X = 2*fold.Position - point.X
			points[point] = true
		}
	} else {
		for point := range points {
			if point.Y > fold.Position {
				tmp = append(tmp, point)
			}
		}
		for _, point := range tmp {
			delete(points, point)
			point.Y = 2*fold.Position - point.Y
			points[point] = true
		}
	}
	return points
}

func findMax(points pointMap) (int, int) {
	var maxX, maxY int

	for point := range points {
		if point.X > maxX {
			maxX = point.X
		}

		if point.Y > maxY {
			maxY = point.Y
		}
	}

	return maxX, maxY
}

func outputMap(points pointMap) string {
	maxX, maxY := findMax(points)
	var str strings.Builder

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			tmp := XYCoords{X: x, Y: y}

			if points[tmp] {
				str.WriteRune('#')
			} else {
				str.WriteRune('.')
			}
		}
		str.WriteRune('\n')
	}

	return str.String()
}

func part1(points pointMap, folds []fold) int {
	return len(doFold(points, folds[0]))
}

func part2(points pointMap, folds []fold) (result string) {
	for _, fold := range folds {
		points = doFold(points, fold)
	}

	return outputMap(points)
}

func main() {
	inputFile := parseArgs()
	inputData, foldData := parseData(getData(inputFile))

	result1 := part1(inputData, foldData)
	fmt.Printf("Part 1: Visible dots: %d\n", result1)

	result2 := part2(inputData, foldData)
	fmt.Printf("Part 2: activation code:\n%v\n", result2)
}
