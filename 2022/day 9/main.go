package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/set"
)

type Move struct {
	direction string
	count     int
}

func InitMove(d string, c int) *Move {
	a := new(Move)
	a.count = c
	a.direction = d
	return a
}

type Rope struct {
	head          Coords
	tail          Coords
	tailPositions set.Interface
}

type Coords [2]int

func InitRope() *Rope {
	a := new(Rope)
	a.head = [2]int{}
	a.tail = [2]int{}
	a.tailPositions = set.New(set.NonThreadSafe)
	a.tailPositions.Add(encodeCoord(a.tail))
	return a
}

func encodeCoord(pos Coords) string {
	return fmt.Sprintf("(%d,%d)", pos[0], pos[1])
}

func (r *Rope) IsValid() bool {
	distance := (r.head[0]-r.tail[0])*(r.head[0]-r.tail[0]) + (r.head[1]-r.tail[1])*(r.head[1]-r.tail[1])
	return distance <= 2
}

func (r *Rope) MoveRope(m *Move) {
	for i := 0; i < m.count; i++ {
		switch m.direction {
		case "R":
			r.head[0] += 1
		case "L":
			r.head[0] -= 1
		case "U":
			r.head[1] += 1
		case "D":
			r.head[1] -= 1
		}

		if !r.IsValid() {
			switch m.direction {
			case "R":
				r.tail[0] = r.head[0] - 1
				r.tail[1] = r.head[1]
			case "L":
				r.tail[0] = r.head[0] + 1
				r.tail[1] = r.head[1]
			case "U":
				r.tail[0] = r.head[0]
				r.tail[1] = r.head[1] - 1
			case "D":
				r.tail[0] = r.head[0]
				r.tail[1] = r.head[1] + 1
			}
		}

		r.tailPositions.Add(encodeCoord(r.tail))
	}

}

func ParseInput(pathToFile string) *[]*Move {
	f, err := os.Open(pathToFile)
	if err != nil {
		panic(fmt.Sprintf("Unable to open file: %s", pathToFile))
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var result = []*Move{}
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(strings.Trim(line, "\n"), " ")
		c, err := strconv.Atoi(split[1])
		if err != nil {
			panic(fmt.Sprintf("Error while parsing string %s", line))
		}
		m := InitMove(split[0], c)
		result = append(result, m)
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

	moves := ParseInput(inputFileName)

	rope := InitRope()
	for _, m := range *moves {
		rope.MoveRope(m)
	}
	fmt.Println(rope.tailPositions.Size())
}
