package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Coords []int
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
	start := Coords{-1, -1}
	end := Coords{-1, -1}
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
	return a <= 1
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

func hash(s Coords) int {
	return s[0]*100000 + s[1]
}

func findPath(grid Grid, s, e Coords) int {
	distances := map[int]int{}
	distances[hash(s)] = 0
	queue := []Coords{s}
	cur := s

	for len(queue) > 0 {
		_, ok := distances[hash(e)]
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

	return distances[hash(e)]
}

func enqueueIfNeeded(distances map[int]int, p Coords, cur Coords, queue []Coords) []Coords {
	v, ok := distances[hash(p)]
	if !ok || v >= distances[hash(cur)]+1 {
		distances[hash(p)] = distances[hash(cur)] + 1

		found := false
		for i := 0; i < len(queue); i++ {
			if queue[i].Equals(p) {
				found = true
			}
		}

		if !found {
			queue = append(queue, p)
		}
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
	fmt.Println("Start and end: ", start, end)
	steps := findPath(grid, start, end)
	fmt.Println("Mimimal steps count: ", steps)
}
