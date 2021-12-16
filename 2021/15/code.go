package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/RyanCarrier/dijkstra"
)

type XYCoords struct {
	X, Y int
}

type RiskMap map[XYCoords]int

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

func parseData(inputData []string, multiplier int) (*dijkstra.Graph, int) {

	risks := make(RiskMap)

	graph := dijkstra.NewGraph()
	var maxId int

	lenY := len(inputData)
	lenX := len(inputData[0])

	// build initial map
	for y, line := range inputData {
		for x, risk := range line {
			risks[XYCoords{X: x, Y: y}] = int(risk - '0')
			graph.AddVertex(getId(x, y, lenX*multiplier))
		}
	}

	// initial copy
	tmp := copyMap(risks)

	// multiply map right
	for i := 1; i < multiplier; i++ {
		tmp = multiplyMap(tmp, lenX, 0)

		for pos, val := range tmp {
			risks[pos] = val
			graph.AddVertex(getId(pos.X, pos.Y, lenX*multiplier))
		}
	}

	// make a copy
	tmp = copyMap(risks)

	// multiply map down
	for i := 1; i < multiplier; i++ {
		tmp = multiplyMap(tmp, 0, lenY)

		for pos, val := range tmp {
			risks[pos] = val
			graph.AddVertex(getId(pos.X, pos.Y, lenX*multiplier))
		}
	}

	for srcPos := range risks {
		srcId := getId(srcPos.X, srcPos.Y, lenX*multiplier)

		neighbours := getNeighbours(srcPos)
		for _, dstPos := range neighbours {
			if _, ok := risks[dstPos]; ok {
				destId := getId(dstPos.X, dstPos.Y, lenX*multiplier)
				val := risks[dstPos]

				graph.AddArc(srcId, destId, int64(val))

				if destId > maxId {
					maxId = destId
				}
			}
		}
	}

	return graph, maxId
}

func copyMap(source RiskMap) (dest RiskMap) {
	dest = make(RiskMap)

	for pos, val := range source {
		dest[pos] = val
	}

	return
}

func multiplyMap(inputMap RiskMap, xInc int, yInc int) (newRiskMap RiskMap) {
	newRiskMap = make(RiskMap)

	for pos, val := range inputMap {
		val++
		if val > 9 {
			val = 1
		}
		newRiskMap[XYCoords{X: pos.X + xInc, Y: pos.Y + yInc}] = val
	}
	return
}

func getId(x int, y int, lenX int) int {
	return x*lenX + y
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

func main() {
	inputFile := parseArgs()
	inputData := getData(inputFile)

	graph, maxId := parseData(inputData, 1)
	best, _ := graph.Shortest(0, maxId)
	result1 := best.Distance
	fmt.Printf("Part 1: total risk: %d\n", result1)

	graph, maxId = parseData(inputData, 5)
	best, _ = graph.Shortest(0, maxId)
	result2 := best.Distance
	fmt.Printf("Part 2: total risk: %d\n", result2)

}
