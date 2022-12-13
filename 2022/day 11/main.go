package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Operation struct {
	opearationLiteral      string
	operationArgumentLeft  string
	operationArgumentRight string
}

func (op *Operation) execute(oldValue int) int {
	left, right := oldValue, oldValue
	if op.operationArgumentLeft != "old" {
		left, _ = strconv.Atoi(op.operationArgumentLeft)
	}
	if op.operationArgumentRight != "old" {
		right, _ = strconv.Atoi(op.operationArgumentLeft)
	}
	switch op.opearationLiteral {
	case "*":
		return left * right
	case "+":
		return left + right
	case "-":
		return left - right
	case "/":
		return left / right
	}
	panic("Error in operations literal")
}

type Test struct {
	divisibleTestValue int
	SuccessTarget      int
	FailureTatget      int
}

func (t *Test) try(worryLevel int) int {
	if worryLevel%t.divisibleTestValue == 0 {
		return t.SuccessTarget
	}

	return t.FailureTatget
}

type Monkey struct {
	number              string
	items               []int
	operation           *Operation
	test                Test
	inspectedItemsCount int
}

func ParseInput(pathToFile string) []*Monkey {
	f, err := os.Open(pathToFile)
	if err != nil {
		panic(fmt.Sprintf("Unable to open file: %s", pathToFile))
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	result := []*Monkey{}
	var m *Monkey = nil
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "\n")
		if strings.HasPrefix(line, "Monkey ") {
			if m != nil {
				result = append(result, m)
			}

			m = new(Monkey)
			m.items = *new([]int)

			m.number = line[len("Monkey ") : len("Monkey ")+1]
		}
		if strings.HasPrefix(line, "Starting items: ") {
			l := line[len("Starting items: "):]
			for _, w := range strings.Split(l, ", ") {
				v, err := strconv.Atoi(w)
				if err != nil {
					panic(fmt.Sprintf("Error in items parsing %s", w))
				}

				m.items = append(m.items, v)
			}
		}

		if strings.HasPrefix(line, "Operation: new = ") {
			l := line[len("Operation: new = "):]
			parts := strings.Split(l, " ")
			if len(parts) > 3 {
				panic(fmt.Sprintf("Parsing error in operation %s", l))
			}

			op := new(Operation)
			op.operationArgumentLeft = parts[0]
			op.opearationLiteral = parts[1]
			op.operationArgumentRight = parts[2]
			m.operation = op
		}
	}

	if m != nil {
		result = append(result, m)
	}

	return result
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

}
