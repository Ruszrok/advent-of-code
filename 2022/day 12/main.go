package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	UNREACHABLE = -10000
)

type Coords [2]int
type GridRow []rune
type Grid []GridRow

func ParseInput(pathToFile string) (Grid, Coords, Coords) {
	f, err := os.Open(pathToFile)
	if err != nil {
		panic(fmt.Sprintf("Unable to open file: %s", pathToFile))
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	result := Grid{}
	start := Coords{}
	end := Coords{}
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "\n")
		result = append(result, GridRow(line))
		for i := 0; i < len(line); i++ {
			if line[i] == 'S' {
				start[0] = i
				start[1] = len(result) - 1
				result[len(result)-1][i] = 96
			}
			if line[i] == 'E' {
				end[0] = i
				end[1] = len(result) - 1
				result[len(result)-1][i] = 123
			}
		}
	}

	return result, start, end
}

func (c Coords) Equals(c1 Coords) bool {
	return c[0] == c1[0] && c[1] == c1[1]
}

func (g Grid) value(p Coords) rune {
	return g[p[1]][p[0]]
}

func canMove(a rune) bool {
	return a <= 1 && a >= 0
}

func (g Grid) CanLeftFromPoint(p Coords) bool {
	if p[0] == 0 {
		return false
	}

	val := g.value(Coords{p[0] - 1, p[1]}) - g.value(p)
	return canMove(val)
}

func (g Grid) CanRightFromPoint(p Coords) bool {
	if p[0] == len(g[0])-1 {
		return false
	}

	val := g.value(Coords{p[0] + 1, p[1]}) - g.value(p)
	return canMove(val)
}

func (g Grid) CanDownFromPoint(p Coords) bool {
	if p[1] == 0 {
		return false
	}

	val := g.value(Coords{p[0], p[1] - 1}) - g.value(p)
	return canMove(val)
}

func (g Grid) CanUpFromPoint(p Coords) bool {
	if p[1] == len(g)-1 {
		return false
	}

	val := g.value(Coords{p[0], p[1] + 1}) - g.value(p)
	return canMove(val)
}

func findPath(grid Grid, s, e Coords) int {
	if s.Equals(e) {
		return 0
	}

	res := [4]int{UNREACHABLE, UNREACHABLE, UNREACHABLE, UNREACHABLE}

	if grid.CanDownFromPoint(s) {
		res[0] = 1 + findPath(grid, Coords{s[0], s[1] - 1}, e)
	}

	if grid.CanUpFromPoint(s) {
		res[1] = 1 + findPath(grid, Coords{s[0], s[1] + 1}, e)
	}

	if grid.CanLeftFromPoint(s) {
		res[2] = 1 + findPath(grid, Coords{s[0] - 1, s[1]}, e)
	}

	if grid.CanRightFromPoint(s) {
		res[3] = 1 + findPath(grid, Coords{s[0] + 1, s[1]}, e)
	}

	fmt.Println(res, s)
	return smartSmallest(res)
}

func smartSmallest(res [4]int) int {
	smallestNumber := UNREACHABLE

	for _, e := range res {
		if smallestNumber == UNREACHABLE {
			smallestNumber = e
		}

		if e != UNREACHABLE && e < smallestNumber {
			smallestNumber = e
		}
	}

	return smallestNumber
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
	fmt.Println(grid)
	steps := findPath(grid, start, end)
	fmt.Println("Mimimal steps count: ", steps)
}
