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
		log.Fatalln(err)
	}

	fileScanner := bufio.NewScanner(fh)

	for fileScanner.Scan() {
		inputData = append(inputData, fileScanner.Text())
	}

	return
}

func getFishes(inputData string) (fishes []int) {

	tmp := strings.Split(inputData, ",")

	for i := 0; i < len(tmp); i++ {
		fish, _ := strconv.Atoi(tmp[i])
		fishes = append(fishes, fish)
	}

	return
}

func goFish(fishes []int, days int) (result int) {
	totalFishesPerDay := make([]int, 9)

	// count initial fish
	for _, day := range fishes {
		totalFishesPerDay[day] += 1
	}

	for day := 0; day < days; day++ {
		newFishes := 0
		for age, fishcount := range totalFishesPerDay {
			if age == 0 {
				// all fishes on day 0 spawn an additional fish (day 8)
				// remove them from day 0 and add them to day 6
				newFishes += fishcount
				totalFishesPerDay[0] -= fishcount
			} else {
				// add all fish of this day to the next day
				// remove them from today
				totalFishesPerDay[age-1] += fishcount
				totalFishesPerDay[age] -= fishcount
			}
		}

		// add the new fishes
		totalFishesPerDay[8] += newFishes
		totalFishesPerDay[6] += newFishes

	}

	for day := 0; day <= 8; day++ {
		result += totalFishesPerDay[day]
	}

	return result
}
func main() {
	inputFile := parseArgs()
	inputData := getData(inputFile)

	fishes := getFishes(inputData[0])

	result1 := goFish(fishes, 80)
	fmt.Printf("Part 1: amount of fishes after 80 days: %d\n", result1)

	result2 := goFish(fishes, 256)
	fmt.Printf("Part 1: amount of fishes after 256 days: %d\n", result2)
}
