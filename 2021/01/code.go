package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func parseArgs() (inputFile string) {
	if len(os.Args) == 2 {
		inputFile = os.Args[1]
	} else {
		log.Fatalf("Usage: %v INPUTFILE", os.Args[0])
	}

	return
}

func main() {
	inputFile := parseArgs()
	fh, err := os.Open(inputFile)

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(fh)

	var vals []int

	for fileScanner.Scan() {
		val, _ := strconv.Atoi(fileScanner.Text())
		vals = append(vals, val)
	}

	fmt.Printf("Part 1: %v measurements are larger\n", part1(vals))
	fmt.Printf("Part 2: %v measurements are larger\n", part2(vals))
}

func part2(vals []int) (depthIncrease int) {
	depthIncrease = 0

	var windows []int

	for idx := 1; idx < len(vals)-1; idx++ {
		windows = append(windows, vals[idx-1]+vals[idx]+vals[idx+1])
	}

	depthIncrease = part1(windows)

	return
}

func part1(vals []int) (depthIncrease int) {
	depthIncrease = 0

	prevVal := 0

	for _, val := range vals {
		if prevVal == 0 {
			fmt.Printf("%v (N/A)\n", val)
		} else {
			if val == prevVal {
				fmt.Printf("%v (no change)\n", val)
			} else if val > prevVal {
				fmt.Printf("%v (increased)\n", val)
				depthIncrease += 1
			} else {
				fmt.Printf("%v (decreased)\n", val)
			}
		}
		prevVal = val
	}

	return
}
