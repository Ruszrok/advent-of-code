package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseInput(pathToFile string) [][]rune {
	f, err := os.Open(pathToFile)
	if err != nil {
		panic(fmt.Sprintf("Unable to open file: %s", pathToFile))
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var result [][]rune
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "\n")
		result = append(result, []rune(line))
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
	//var answer2 = 0
	var usedRunes = map[int]bool{}

	//game 1
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			code := encode(i, j)
			if !usedRunes[code] {
				ch := input[i][j]
				if isNumber(ch) {
					usedRunes[code] = true
					numString := string(ch)
					k := j + 1
					for ; k < len(input[i]); k++ {
						ch2 := input[i][k]
						if !isNumber(ch2) {
							break
						}
						numString += string(ch2)
						usedRunes[encode(i, k)] = true
					}
					println(numString)
					if isPartNumber(input, i, j, k-1) {
						println("*")

						v, err := strconv.Atoi(numString)
						if err != nil {
							panic(fmt.Sprintf("Error while parsing string %s", numString))
						}
						answer1 += v
					}
				}
			}
		}
	}

	////game 2
	//for _, s := range input {
	//	maxState := [3]int{0, 0, 0}
	//	for _, nums := range s {
	//		maxState[0] = max(maxState[0], nums[0])
	//		maxState[1] = max(maxState[1], nums[1])
	//		maxState[2] = max(maxState[2], nums[2])
	//	}
	//
	//	answer2 += maxState[0] * maxState[1] * maxState[2]
	//}

	fmt.Println("Sum of games: ", answer1, 4361)
	//fmt.Println("Sum of games: ", answer2, 59795)
}

func isPartNumber(in [][]rune, i, start, end int) bool {
	s := start
	e := end
	if start > 0 {
		s = start - 1
	}

	if end < len(in[i])-1 {
		e = end + 1
	}

	if i == 4 {
		println("s", s, "e", e)
	}

	if isNotNumberOrDot(in[i][s]) {
		return true
	}

	if isNotNumberOrDot(in[i][e]) {
		return true
	}

	if i > 0 {
		for j := s; j <= e; j++ {
			if isNotNumberOrDot(in[i-1][j]) {
				return true
			}
		}
	}

	if i < len(in)-1 {
		for j := s; j <= e; j++ {
			if isNotNumberOrDot(in[i+1][j]) {
				return true
			}
		}
	}

	return false
}

func isNotNumberOrDot(ch rune) bool {
	return !isNumber(ch) && ch != '.'
}

func isNumber(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func encode(i, j int) int {
	return i*1000 + j
}
