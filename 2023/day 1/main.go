package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func ParseInput(pathToFile string) []string {
	f, err := os.Open(pathToFile)
	if err != nil {
		panic(fmt.Sprintf("Unable to open file: %s", pathToFile))
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var result []string
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "\n")
		result = append(result, line)
	}

	return result
}

var numberMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func getNumberFromSpelling(s string, index int) (bool, int) {
	for i := 5; i >= 3; i-- {
		if index+i <= len(s) {
			v, exists := numberMap[s[index:index+i]]
			if exists {
				return true, v
			}
		}
	}

	return false, 0
}

func main() {
	isTestFile := false
	isTest2File := false
	flag.BoolVar(&isTestFile, "t", false, "display in uppercase")
	flag.BoolVar(&isTest2File, "tt", false, "display in uppercase")
	flag.Parse()
	inputFileName := "test.txt"
	if isTest2File {
		inputFileName = "test2.txt"
	}
	if !isTestFile && !isTest2File {
		inputFileName = "input.txt"
	}

	input := ParseInput(inputFileName)
	answer := 0

	for _, s := range input {
		var nums []int
		for i := 0; i < len(s); i++ {
			ch := s[i]
			if ch >= '0' && ch <= '9' {
				nums = append(nums, int(ch-'0'))
			}

			var m, val = getNumberFromSpelling(s, i)
			if m {
				nums = append(nums, val)
			}
		}
		answer += nums[0]*10 + nums[len(nums)-1]
	}

	//fmt.Println("Sum of codes: ", answer, 55090)
	fmt.Println("Sum of codes: ", answer, 54845)
}
