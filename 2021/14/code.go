package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type polymer map[string]int
type polymerRules map[string][]string

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

func parseData(inputData []string) (template polymer, rules polymerRules) {

	rules = make(polymerRules)
	template = make(polymer)

	var a, b, c rune

	for _, line := range inputData {
		if len(line) == 0 {
			continue
		} else if strings.Contains(line, "->") {
			// rule
			cnt, err := fmt.Sscanf(line, "%c%c -> %c", &a, &b, &c)

			if err != nil {
				log.Fatalln(err)
			}

			if cnt != 3 {
				log.Fatalln("Unexpected string")
			}

			origPolymer := string(a) + string(b)
			polymerA := string(a) + string(c)
			polymerB := string(c) + string(b)

			rules[origPolymer] = make([]string, 2)
			rules[origPolymer][0] = polymerA
			rules[origPolymer][1] = polymerB

		} else {
			// start template
			for idx := 0; idx < len(line)-1; idx++ {
				template = insertItem(template, line[idx:idx+2])
			}
			template = insertItem(template, string(line[0]))
			template = insertItem(template, string(line[len(line)-1]))
		}
	}

	return template, rules
}

func insertItem(template polymer, itemToInsert string) polymer {
	template[itemToInsert]++

	return template
}

func runRules(template polymer, rules polymerRules) polymer {
	newTemplate := make(polymer)

	for item := range template {
		for _, rule := range rules[item] {
			newTemplate[rule] += template[item]
		}

		if len(item) == 1 {
			newTemplate[item]++
		}
	}
	return newTemplate
}

func countElements(template polymer) []int {
	tmp := make(map[rune]int)

	for item := range template {
		for _, char := range item {
			tmp[char] += template[item]
		}
	}

	var cnt []int
	for _, i := range tmp {
		cnt = append(cnt, i/2)
	}
	sort.Ints(cnt)

	return cnt
}

func process(template polymer, rules polymerRules, steps int) int {
	for i := 0; i < steps; i++ {
		template = runRules(template, rules)
	}

	cnt := countElements(template)

	return cnt[len(cnt)-1] - cnt[0]

}

func main() {
	inputFile := parseArgs()
	template, rules := parseData(getData(inputFile))

	result1 := process(template, rules, 10)
	fmt.Printf("Part 1: %d steps, diff: %d\n", 10, result1)

	result2 := process(template, rules, 40)
	fmt.Printf("Part 2: %d steps, diff: %d\n", 40, result2)
}
