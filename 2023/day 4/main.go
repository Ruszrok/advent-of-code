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

			println(aN[j], wN[winPosition])
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
	}

	////game 2
	//for i := 0; i < len(input); i++ {
	//	for j := 0; j < len(input[i]); j++ {
	//		if input[i][j] == '*' {
	//			var closeParts = getClosesParts(partPositions, i, j)
	//			if len(closeParts) == 2 {
	//				answer2 += closeParts[0].value * closeParts[1].value
	//			}
	//		}
	//	}
	//}

	fmt.Println("Sum of games: ", answer1, 13)
	fmt.Println("Sum of games: ", answer2, 467835)
}
func pow2(n int) int {
	a := 1
	for i := 0; i < n; i++ {
		a *= 2
	}
	return a
}
