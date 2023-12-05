package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Card struct {
	winNumbers    []int
	actualNumbers []int
}

func newCard(w, a []int) *Card {
	p := Card{winNumbers: w, actualNumbers: a}
	return &p
}
func ParseInput(pathToFile string) []*Card {
	f, err := os.Open(pathToFile)
	if err != nil {
		panic(fmt.Sprintf("Unable to open file: %s", pathToFile))
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var result []*Card
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "\n")
		colonIndex := strings.Index(line, ":")
		line = line[colonIndex+1:]
		line = strings.Trim(line, " ")
		split := strings.Split(line, "|")
		var winNumbers []int
		winSplit := strings.Split(split[0], " ")
		for _, s := range winSplit {
			if s == "" {
				continue
			}
			s = strings.Trim(s, " ")
			v, err := strconv.Atoi(s)
			if err != nil {
				panic(fmt.Sprintf("Error while parsing string %s", s))
			}
			winNumbers = append(winNumbers, v)
		}
		var actualNumber []int
		actualSplit := strings.Split(split[1], " ")
		for _, s := range actualSplit {
			if s == "" {
				continue
			}
			s = strings.Trim(s, " ")
			v, err := strconv.Atoi(s)
			if err != nil {
				panic(fmt.Sprintf("Error while parsing string %s", s))
			}
			actualNumber = append(actualNumber, v)
		}

		result = append(result, newCard(winNumbers, actualNumber))
	}

	return result
}

func main() {
	isTestFile := false
	isTest2File := false
	flag.BoolVar(&isTestFile, "t", false, "display in uppercase")
	flag.BoolVar(&isTest2File, "tt", false, "display in uppercase")
	flag.Parse()
	inputFileName := "test.txt"
	if isTest2File {
		inputFileName = "test2.txt"
	}
	if !isTestFile && !isTest2File {
		inputFileName = "input.txt"
	}

	input := ParseInput(inputFileName)
	var answer1 = 0
	var answer2 = 0

	var matchesCount []int
	//game 1
	for i := 0; i < len(input); i++ {
		wN := input[i].winNumbers
		aN := input[i].actualNumbers
		sort.Slice(wN, func(i, j int) bool {
			return wN[i] < wN[j]
		})
		sort.Slice(aN, func(i, j int) bool {
			return aN[i] < aN[j]
		})

		winPosition := 0
		matches := 0
		for j := 0; j < len(aN); j++ {
			if winPosition == len(wN) {
				break
			}

			if aN[j] > wN[winPosition] {
				winPosition++
				j--
				continue
			}

			if aN[j] == wN[winPosition] {
				matches++
			}
		}
		if matches > 0 {
			answer1 += pow2(matches - 1)
		}

		matchesCount = append(matchesCount, matches)
	}

	scratches := make([]int, len(matchesCount))
	//game 2
	for i := 0; i < len(matchesCount); i++ {
		scratches[i]++
		if matchesCount[i] > 0 {
			for j := i + 1; j <= i+matchesCount[i]; j++ {
				scratches[j] += scratches[i]
			}
		}
	}
	for _, s := range scratches {
		answer2 += s
	}

	fmt.Println("Sum of games: ", answer1, 21138)
	fmt.Println("Sum of games: ", answer2, 30)
}

func pow2(n int) int {
	a := 1
	for i := 0; i < n; i++ {
		a *= 2
	}
	return a
}
