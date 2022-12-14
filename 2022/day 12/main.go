package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	UNREACHABLE = 100000000000
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
	return a == 0 || a == 1
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

func dequeue(o []Coords) (Coords, []Coords) {
	if len(o) == 1 {
		return o[0], []Coords{}
	}

	return o[0], o[1:]
}

func findPath(grid Grid, s, e Coords) int {
	distances := map[Coords]int{}
	distances[s] = 0
	queue := []Coords{s}
	cur := s

	for len(queue) > 0 {
		_, ok := distances[e]
		if ok {
			break
		}
		cur, queue = dequeue(queue)
		if grid.CanDownFromPoint(cur) {
			p := Coords{cur[0], cur[1] - 1}
			queue = enqueueIfNeeded(distances, p, cur, queue)
		}

		if grid.CanUpFromPoint(cur) {
			p := Coords{cur[0], cur[1] + 1}
			queue = enqueueIfNeeded(distances, p, cur, queue)
		}

		if grid.CanLeftFromPoint(cur) {
			p := Coords{cur[0] - 1, cur[1]}
			queue = enqueueIfNeeded(distances, p, cur, queue)
		}

		if grid.CanRightFromPoint(cur) {
			p := Coords{cur[0] + 1, cur[1]}
			queue = enqueueIfNeeded(distances, p, cur, queue)
		}
	}

	return distances[e]
}

func enqueueIfNeeded(distances map[Coords]int, p Coords, cur Coords, queue []Coords) []Coords {
	v, ok := distances[p]
	if !ok || v >= distances[cur]+1 {
		distances[p] = distances[cur] + 1
		queue = append(queue, p)
	}
	return queue
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
	fmt.Println("Mimimal steps count: ", steps)
}
