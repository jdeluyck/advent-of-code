package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func doCalc(inputData []string, diagonal bool) (result int) {

	breakpoints := make(map[string]int)

	for _, line := range inputData {
		var fromX, fromY, toX, toY int

		cnt, err := fmt.Sscanf(line, "%d,%d -> %d,%d", &fromX, &fromY, &toX, &toY)

		if err != nil {
			log.Fatalln(err)
		}

		if cnt != 4 {
			log.Fatalln("Couldn't parse 4 fields! Danger Danger!")
		}

		tmp := make([]string, 0)

		// straight lines, we like straight lines
		if fromX == toX || fromY == toY {
			// Swap stuff around so we only need to add up. We're lazy.
			if fromY > toY {
				tmpY := fromY
				fromY = toY
				toY = tmpY
			}
			if fromX > toX {
				tmpX := fromX
				fromX = toX
				toX = tmpX
			}

			for x := fromX; x <= toX; x++ {
				for y := fromY; y <= toY; y++ {
					tmp = append(tmp, fmt.Sprintf("%d,%d", x, y))
				}
			}
		} else {
			// 45 degree crap
			if diagonal {
				// Diagonals can go two ways, make sure we catch that
				modX := 1
				modY := 1

				if fromX > toX {
					modX = -1
				}
				if fromY > toY {
					modY = -1
				}

				for diff := 0; diff < ((toX-fromX)*modX)+1; diff++ {
					posX := fromX + (diff * modX)
					posY := fromY + (diff * modY)
					tmp = append(tmp, fmt.Sprintf("%d,%d", posX, posY))
				}
			}
		}

		// iterate over points to add stuff
		for _, point := range tmp {
			_, keyExists := breakpoints[point]

			if keyExists {
				// key exists, add 1 to it
				breakpoints[point] += 1
			} else {
				// key doesn't exist, init it
				breakpoints[point] = 1
			}
		}
	}

	for _, intersects := range breakpoints {
		if intersects > 1 {
			result += 1
		}
	}

	return

}

func main() {
	inputFile := parseArgs()
	inputData := getData(inputFile)

	result1 := doCalc(inputData, false)
	result2 := doCalc(inputData, true)

	fmt.Printf("Part 1: points 2 or more overlap: %v\n", result1)
	fmt.Printf("Part 2: points 2 or more overlap: %v\n", result2)
}
