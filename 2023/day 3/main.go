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

type PartPosition struct {
	row   int
	start int
	end   int
}

func newPartPosition(r, s, e int) *PartPosition {
	p := PartPosition{row: r, start: s, end: e}
	return &p
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
	var partPositions []*PartPosition
	//game 1
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			ch := input[i][j]
			if isNumber(ch) {
				numString := string(ch)
				k := j + 1
				for ; k < len(input[i]); k++ {
					ch2 := input[i][k]
					if !isNumber(ch2) {
						break
					}
					numString += string(ch2)
				}

				pp := newPartPosition(i, j, k-1)
				if isPartNumber(input, pp) {
					partPositions = append(partPositions, pp)
					v, err := strconv.Atoi(numString)
					if err != nil {
						panic(fmt.Sprintf("Error while parsing string %s", numString))
					}
					answer1 += v
				}

				j = k - 1
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

	fmt.Println("Sum of games: ", answer1, 520135)
	//fmt.Println("Sum of games: ", answer2, 59795)
}

func isPartNumber(in [][]rune, pp *PartPosition) bool {
	s := pp.start
	e := pp.end
	if pp.start > 0 {
		s = pp.start - 1
	}

	if pp.end < len(in[pp.row])-1 {
		e = pp.end + 1
	}

	if isNotNumberOrDot(in[pp.row][s]) {
		return true
	}

	if isNotNumberOrDot(in[pp.row][e]) {
		return true
	}

	if pp.row > 0 {
		for j := s; j <= e; j++ {
			if isNotNumberOrDot(in[pp.row-1][j]) {
				return true
			}
		}
	}

	if pp.row < len(in)-1 {
		for j := s; j <= e; j++ {
			if isNotNumberOrDot(in[pp.row+1][j]) {
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
