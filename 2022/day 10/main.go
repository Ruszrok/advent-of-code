package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Operations string

const (
	addx Operations = "addx"
	noop Operations = "noop"
)

type Operation struct {
	name Operations
	args int
}

func InitOperation(name Operations, args int) *Operation {
	o := new(Operation)
	o.name = name
	o.args = args
	return o
}

func ParseInput(pathToFile string) []*Operation {
	f, err := os.Open(pathToFile)
	if err != nil {
		panic(fmt.Sprintf("Unable to open file: %s", pathToFile))
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var result = []*Operation{}
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "\n")
		if line == string(noop) {
			result = append(result, InitOperation(noop, 0))
		} else {
			split := strings.Split(line, " ")
			c, err := strconv.Atoi(split[1])
			if err != nil {
				panic(fmt.Sprintf("Error while parsing string %s", line))
			}
			result = append(result, InitOperation(addx, c))
		}
	}

	return result
}

func dequeue(o []*Operation) (*Operation, []*Operation) {
	if len(o) == 1 {
		return o[0], []*Operation{}
	}

	return o[0], o[1:]
}

func main() {
	isTestFile := false
	flag.BoolVar(&isTestFile, "t", false, "display in uppercase")
	flag.Parse()
	inputFileName := "test.txt"
	if !isTestFile {
		inputFileName = "input.txt"
	}

	operations := ParseInput(inputFileName)

	op_costs := make(map[Operations]int)
	op_costs[noop] = 1
	op_costs[addx] = 2

	sum := 0
	next_control := 20
	x := 1
	cycleCount := 1
	currentOpCost := 1
	current := InitOperation(noop, 0)
	// screen := strings.Repeat(".", 240)
	screen := [240]byte{}
	for i := 0; i < len(screen); i++ {
		screen[i] = '.'
	}

	for cycleCount <= 240 && len(operations) > 0 {
		currentOpCost -= 1
		if currentOpCost == 0 {
			switch current.name {
			case addx:
				x += current.args
			}

			current, operations = dequeue(operations)
			switch current.name {
			case noop:
				currentOpCost = op_costs[noop]
			case addx:
				currentOpCost = op_costs[addx]
			}
		}

		if cycleCount == next_control {
			sum += cycleCount * x
			next_control += 40
		}

		mod := cycleCount % 40
		if (x <= mod && mod <= x+2) || (mod == 0 && x >= 38) {
			screen[cycleCount-1] = '#'
		}
		cycleCount += 1
	}

	fmt.Println("Signal strength: ", sum)
	for i := 0; i < 240; i = i + 40 {
		fmt.Println(string(screen[i : i+40]))
	}
}
