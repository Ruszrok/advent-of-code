package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
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
		right, _ = strconv.Atoi(op.operationArgumentRight)
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
	ExecutionCount     int
}

func (t *Test) try(worryLevel int) int {
	t.ExecutionCount++
	if worryLevel%t.divisibleTestValue == 0 {
		return t.SuccessTarget
	}

	return t.FailureTatget
}

type Monkey struct {
	number    string
	items     []int
	operation *Operation
	test      *Test
}

func PreconditionTest(m *Monkey) {
	if m == nil || m.items == nil || m.operation == nil || m.test == nil {
		panic("Partial nil is nill")
	}
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
				PreconditionTest(m)
				result = append(result, m)
			}

			m = new(Monkey)
			m.items = *new([]int)

			m.number = line[len("Monkey ") : len("Monkey ")+1]
		}

		if strings.HasPrefix(line, "  Starting items: ") {
			l := line[len("  Starting items: "):]
			for _, w := range strings.Split(l, ", ") {
				v, err := strconv.Atoi(w)
				if err != nil {
					panic(fmt.Sprintf("Error in items parsing %s", w))
				}

				m.items = append(m.items, v)
			}
		}

		if strings.HasPrefix(line, "  Operation: new = ") {
			l := line[len("  Operation: new = "):]
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

		if strings.HasPrefix(line, "  Test: divisible by ") {
			l := line[len("  Test: divisible by "):]
			v, err := strconv.Atoi(l)
			if err != nil {
				panic(fmt.Sprintf("Error in parsing %s", l))
			}

			t := new(Test)
			t.divisibleTestValue = v
			t.ExecutionCount = 0
			m.test = t
		}

		if strings.HasPrefix(line, "    If false: throw to monkey ") {
			l := line[len("    If false: throw to monkey "):]
			v, err := strconv.Atoi(l)
			if err != nil {
				panic(fmt.Sprintf("Error in parsing %s", l))
			}

			m.test.SuccessTarget = v
		}

		if strings.HasPrefix(line, "    If false: throw to monkey ") {
			l := line[len("    If false: throw to monkey "):]
			v, err := strconv.Atoi(l)
			if err != nil {
				panic(fmt.Sprintf("Error in parsing %s", l))
			}

			m.test.FailureTatget = v
		}
	}

	if m != nil {
		PreconditionTest(m)
		result = append(result, m)
	}

	return result
}

func dequeue(o []int) (int, []int) {
	if len(o) == 1 {
		return o[0], []int{}
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

	monkeys := ParseInput(inputFileName)

	maxCount := 2
	for i := 0; i < maxCount; i++ {
		fmt.Printf("State before operation %d\n", i)
		for i := 0; i < len(monkeys); i++ {
			fmt.Println(monkeys[i].items)
		}
		for j := 0; j < len(monkeys); j++ {
			m := monkeys[j]
			item := -1
			for len(m.items) > 0 {
				item, m.items = dequeue(m.items)
				nextlevel := m.operation.execute(item)
				nextlevel /= 3
				nextMonkey := m.test.try(nextlevel)
				monkeys[nextMonkey].items = append(monkeys[nextMonkey].items, nextlevel)
			}
		}
	}

	for i := 0; i < len(monkeys); i++ {
		fmt.Println(monkeys[i].test.ExecutionCount)
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].test.ExecutionCount > monkeys[j].test.ExecutionCount
	})

	power := monkeys[0].test.ExecutionCount * monkeys[1].test.ExecutionCount
	fmt.Println(power)
}
