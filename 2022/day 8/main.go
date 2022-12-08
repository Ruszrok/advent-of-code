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

		for i, c := range strings.Trim(line, "\n") {
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
	if visible {
		return 1
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

func GetScenicScore(forest *[][]int, posX, posY, rowC, colC int) int {
	if posX == 0 || posX == rowC-1 || posY == 0 || posY == colC-1 {
		return 0
	}

	f := *forest
	height := f[posX][posY]

	left := 0
	for i := posY - 1; i >= 0; i-- {
		left += 1
		if f[posX][i] >= height {
			break
		}
	}

	right := 0
	for i := posY + 1; i < colC; i++ {
		right += 1
		if f[posX][i] >= height {
			break
		}
	}

	top := 0
	for i := posX - 1; i >= 0; i-- {
		top += 1
		if f[i][posY] >= height {
			break
		}
	}

	down := 0
	for i := posX + 1; i < rowC; i++ {
		down += 1
		if f[i][posY] >= height {
			break
		}
	}

	return left * right * top * down
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
	rowsC, columnsC := len(*parsedInput), len((*parsedInput)[0])

	count := 0
	maxScore := 0
	for i := 0; i < rowsC; i++ {
		for j := 0; j < columnsC; j++ {
			visibility := IsVisible(parsedInput, i, j, rowsC, columnsC)
			score := GetScenicScore(parsedInput, i, j, rowsC, columnsC)
			count += visibility
			if maxScore < score {
				maxScore = score
			}
		}
	}

	fmt.Println(count, "Expected: ", 1851)
	fmt.Println(maxScore, "Expected: ", 574080)
}
