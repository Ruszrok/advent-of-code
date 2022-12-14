package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Coords [2]int
type GridRow []rune

func ParseInput(pathToFile string) ([]GridRow, Coords, Coords) {
	f, err := os.Open(pathToFile)
	if err != nil {
		panic(fmt.Sprintf("Unable to open file: %s", pathToFile))
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	result := []GridRow{}
	start := Coords{}
	end := Coords{}
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "\n")
		result = append(result, GridRow(line))
		for i := 0; i < len(line); i++ {
			if line[i] == 'S' {
				start[0] = i
				start[1] = len(result) - 1
			}
			if line[i] == 'E' {
				end[0] = i
				end[1] = len(result) - 1
			}
		}
	}

	return result, start, end
}

func findPath(grid []GridRow, s, e Coords) int {
	return 0
}

func main() {
	isTestFile := false
	flag.BoolVar(&isTestFile, "t", false, "display in uppercase")
	flag.Parse()
	inputFileName := "test.txt"
	if !isTestFile {
		inputFileName = "input.txt"
	}

	grid, start, end := ParseInput(inputFileName)

	steps := findPath(grid, start, end)
	fmt.Println("Mimimal steps count: ", steps, start, end)
}
