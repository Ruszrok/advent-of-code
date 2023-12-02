package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	NOT_A_VALUE = -40000000
)

type CmpPair struct {
	left  string
	right string
}

func ParseInput(pathToFile string) []*CmpPair {
	f, err := os.Open(pathToFile)
	if err != nil {
		panic(fmt.Sprintf("Unable to open file: %s", pathToFile))
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	result := []*CmpPair{}
	cur := new(CmpPair) 
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "\n")
		if line == "" {
			result = append(result, cur)
			cur = new(CmpPair)
			continue
		}
		if cur.left == "" {
			cur.left = line
		}
		if cur.right == "" {
			cur.right = line
		}
	}
	result = append(result, cur)

	return result
}

func IsRightOrder(p *CmpPair) bool {
	l := NewComplexList(p.left)
	r := NewComplexList(p.right)

	isRight := true
	curL := l.next()
	for curL != NOT_A_VALUE {
		curR := r.next()
		if curR == NOT_A_VALUE {
			isRight = false
			break
		}
		if curR > curL
	}
	return isRight
}

type ComplexList struct {
	values []ComplexList
}

type ComplexListElement struct {
	isList bool
	val    int
	list   *ComplexList
}

func NewComplexList(s string) *ComplexList {
	return nil
}

func (c *ComplexList) next() int {
	return NOT_A_VALUE
}

func main() {
	isTestFile := false
	flag.BoolVar(&isTestFile, "t", false, "display in uppercase")
	flag.Parse()
	inputFileName := "test.txt"
	if !isTestFile {
		inputFileName = "input.txt"
	}

	pairs := ParseInput(inputFileName)
	answer := 0

	for i, p := range pairs {
		if IsRightOrder(p) {
			answer += i + 1
		}
	}

	fmt.Println("Sum of right order indexes: ", answer)
}
