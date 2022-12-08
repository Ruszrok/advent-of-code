package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseInput(pathToFile string) *[][]int {
	f, err := os.Open(pathToFile)
	if err != nil {
		panic(fmt.Sprintf("Unable to open file: %s", pathToFile))
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var result = [][]int{}
	lineNumber := 0
	for scanner.Scan() {
		result = append(result, []int{})
		line := scanner.Text()
		parsedStr := strings.Trim(line, "\n")
		for i, c := range parsedStr {
			v, err := strconv.Atoi(string(c))
			if err != nil {
				panic(fmt.Sprintf("Parsing error line: %s | position: %d | rune: %c", line, i, c))
			}

			result[lineNumber] = append(result[lineNumber], v)
		}
		lineNumber += 1
	}
	return &result
}

func main() {
	isTestFile := false
	flag.BoolVar(&isTestFile, "t", false, "display in uppercase")
	flag.Parse()
	inputFileName := "test.txt"
	if !isTestFile {
		inputFileName = "input.txt"
	}
	parsed := ParseInput(inputFileName)
	for _, v := range *parsed {
		for _, h := range v {
			fmt.Printf("%d", h)
		}
		fmt.Printf("\n")
	}

}
