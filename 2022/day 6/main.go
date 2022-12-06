package main

import (
	"bufio"
	"fmt"
	"os"
)

type queue []rune

func (q queue) pop() (queue, rune) {
	return q[1:], q[0]
}

func (q queue) contains(param rune) bool {
	found := false
	for _, c := range q {
		if c == param {
			found = true
			break
		}
	}
	return found
}

func ParseInput(pathToFile string) string {
	f, err := os.Open(pathToFile)
	if err != nil {
		panic(fmt.Sprintf("Unable to open file: %s", pathToFile))
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		return scanner.Text()
	}

	return ""
}

func GetCodePostion(input string, size int) int {
	var q queue

	for i, c := range input {
		if len(q) == size {
			return i
		}

		if !q.contains(c) {
			q = append(q, c)
		} else {
			el := '*'
			for el != c {
				q, el = q.pop()
			}
			q = append(q, c)
		}
	}
	return -1
}

func main() {
	var s = "input.txt"
	input := ParseInput(s)
	res := GetCodePostion(input, 14)
	fmt.Println(res)
}
