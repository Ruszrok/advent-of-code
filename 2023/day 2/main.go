package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseInput(pathToFile string) [][][3]int {
	f, err := os.Open(pathToFile)
	if err != nil {
		panic(fmt.Sprintf("Unable to open file: %s", pathToFile))
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var result [][][3]int
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "\n")
		//index of : in line
		colonIndex := strings.Index(line, ":")
		line = line[colonIndex+1:]
		line = strings.Trim(line, " ")
		split := strings.Split(line, ";")
		result = append(result, [][3]int{})

		for _, s := range split {
			s = strings.Trim(s, " ")
			split2 := strings.Split(s, ",")
			var nums [3]int
			for _, s2 := range split2 {
				s2 = strings.Trim(s2, " ")
				s3 := strings.Split(s2, " ")
				nums[col[s3[1]]], err = strconv.Atoi(s3[0])
				if err != nil {
					panic(fmt.Sprintf("Error while parsing string %s", s2))
				}
			}

			result[len(result)-1] = append(result[len(result)-1], nums)
		}
	}

	return result
}

var col = map[string]int{
	"red":   0,
	"green": 1,
	"blue":  2,
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
	var cubeCombination = [3]int{}
	cubeCombination[col["red"]] = 12
	cubeCombination[col["green"]] = 13
	cubeCombination[col["blue"]] = 14

	//game 1
	for i, s := range input {
		gamePossible := true
		for _, nums := range s {
			if nums[0] > cubeCombination[0] || nums[1] > cubeCombination[1] || nums[2] > cubeCombination[2] {
				gamePossible = false
				break
			}
		}

		if gamePossible {
			answer1 += i + 1
		}
	}

	//game 2
	for _, s := range input {
		maxState := [3]int{0, 0, 0}
		for _, nums := range s {
			maxState[0] = max(maxState[0], nums[0])
			maxState[1] = max(maxState[1], nums[1])
			maxState[2] = max(maxState[2], nums[2])
		}

		answer2 += maxState[0] * maxState[1] * maxState[2]
	}

	fmt.Println("Sum of games: ", answer1, 2617)
	fmt.Println("Sum of games: ", answer2, 59795)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
