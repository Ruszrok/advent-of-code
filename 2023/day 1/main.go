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
	result := []string{}
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "\n")
		result = append(result, line)
	}

	return result
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
		nums := []int{}
		for _, ch := range s {
			if ch >= '0' && ch <= '9' {
				nums = append(nums, int(ch-'0'))
			}
		}
		answer += nums[0]*10 + nums[len(nums)-1]
	}

	fmt.Println("Sum of codes: ", answer, 55090)
}
