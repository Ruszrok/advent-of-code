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

func InitializeVisibilityMap(forest *[][]int) (*[][]int, int, int) {
	numRows := len(*forest)
	numColumns := len((*forest)[0])
	visibilityMap := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		visibilityMap[i] = make([]int, numColumns)
	}
	return &visibilityMap, numRows, numColumns
}

func IsVisible(forest *[][]int, posX, posY, rowC, colC int) int {
	if posX == 0 || posX == rowC-1 || posY == 0 || posY == colC-1 {
		return 1
	}

	f := *forest
	height := f[posX][posY]
	visible := true
	for i := posY - 1; i >= 0; i-- {
		if f[posX][i] >= height {
			visible = false
			break
		}
	}
	if visible {
		return 1
	}
	visible = true
	for i := posY + 1; i < colC; i++ {
		if f[posX][i] >= height {
			visible = false
			break
		}
	}

	visible = true
	for i := posX - 1; i >= 0; i-- {
		if f[i][posY] >= height {
			visible = false
			break
		}
	}
	if visible {
		return 1
	}

	visible = true
	for i := posX + 1; i < rowC; i++ {
		if f[i][posY] >= height {
			visible = false
			break
		}
	}
	if visible {
		return 1
	}

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
	parsedInput := ParseInput(inputFileName)

	// Initialize visibility map
	visibilityMap, rowsC, columnsC := InitializeVisibilityMap(parsedInput)

	count := 0
	for i := 0; i < rowsC; i++ {
		for j := 0; j < columnsC; j++ {
			(*visibilityMap)[i][j] = IsVisible(parsedInput, i, j, rowsC, columnsC)
			count += (*visibilityMap)[i][j]
		}
	}
	fmt.Println(*visibilityMap)
	fmt.Println(count)
}
