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

type Coords [2]int

func encodeCoord(pos Coords) string {
	return fmt.Sprintf("(%d,%d)", pos[0], pos[1])
}

type MultiKnotRope struct {
	knotCount     int
	knots         []Coords
	tailPositions set.Interface
}

func InitMultiKnotRope(size int) *MultiKnotRope {
	r := new(MultiKnotRope)
	r.knotCount = size
	r.knots = make([]Coords, size)
	r.tailPositions = set.New(set.NonThreadSafe)
	r.tailPositions.Add(encodeCoord(r.knots[r.knotCount-1]))
	return r
}

func (r *MultiKnotRope) MoveRope(m *Move) {
	for i := 0; i < m.count; i++ {
		r.oneTimeMove(m.direction)
		r.tailPositions.Add(encodeCoord(r.knots[r.knotCount-1]))
	}
}

func (r *MultiKnotRope) oneTimeMove(direction string) {
	//head always moves to direction
	movePointLinear(&r.knots[0], direction)
	if r.IsValid() {
		return
	}

	for i := 1; i < r.knotCount; i++ {
		if distanceSquare(r.knots[i-1], r.knots[i]) > 4 {
			movePointDiagonaly(&r.knots[i], r.knots[i-1], direction)
		} else {
			proximityMoveLinear(&r.knots[i], r.knots[i-1])
		}

		if r.IsValid() {
			break
		}
	}

	if !r.IsValid() {
		fmt.Println(r.knots)
		panic("Rope was broken on OneTimeMove")
	}
}

func movePointDiagonaly(p *Coords, target Coords, direction string) {
	diags := [4]Coords{{p[0] + 1, p[1] + 1}, {p[0] - 1, p[1] + 1}, {p[0] - 1, p[1] - 1}, {p[0] + 1, p[1] - 1}}
	for _, d := range diags {
		if distanceSquare(target, d) <= 2 {
			p[0] = d[0]
			p[1] = d[1]
			return
		}
	}

	panic(fmt.Sprintf("Diagonlal Move failed Knot = %v, Target=%v, Direction=%s", p, target, direction))
}

func proximityMoveLinear(p *Coords, t Coords) {
	if (p[0]-t[0])*(p[0]-t[0]) == 4 {
		p[0] = p[0] - (p[0]-t[0])/2
	} else {
		p[1] = p[1] - (p[1]-t[1])/2
	}
}

func movePointLinear(p *Coords, direction string) {
	switch direction {
	case "R":
		p[0] += 1
	case "L":
		p[0] -= 1
	case "U":
		p[1] += 1
	case "D":
		p[1] -= 1
	default:
		panic(fmt.Sprintf("Unknown direction: %s", direction))
	}
}

func (r *MultiKnotRope) IsValid() bool {
	for i := 1; i < r.knotCount; i++ {
		k1 := r.knots[i]
		k2 := r.knots[i-1]
		if distanceSquare(k1, k2) > 2 {
			return false
		}
	}
	return true
}

func distanceSquare(k1 Coords, k2 Coords) int {
	return (k1[0]-k2[0])*(k1[0]-k2[0]) + (k1[1]-k2[1])*(k1[1]-k2[1])
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
	inputFileName := "test1.txt"
	if !isTestFile {
		inputFileName = "input.txt"
	}

	moves := ParseInput(inputFileName)

	rope := InitMultiKnotRope(2)
	mkRope := InitMultiKnotRope(10)
	for _, m := range *moves {
		rope.MoveRope(m)
		mkRope.MoveRope(m)
	}

	fmt.Println(rope.tailPositions.Size(), "Expected 6090 main, Expected 13 test, Expected 88 test 1")
	fmt.Println(mkRope.tailPositions.Size(), "Expected 2566 main, Expected 1 test, Expected 36 test 1")
}
