package main

import (
	"fmt"
	"testing"
)

func TestGetCodePostion(t *testing.T) {
	tests := []struct {
		input    string
		size     int
		expected int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4, 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 4, 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 4, 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4, 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4, 11},
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14, 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 14, 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 14, 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 14, 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 14, 26},
	}

	for _, ts := range tests {
		t.Run(fmt.Sprintf("Get code for input: %s", ts.input), func(t *testing.T) {
			res := GetCodePostion(ts.input, ts.size)
			if res != ts.expected {
				t.Fatalf("Wrong anser. Expected = %d. Got = %d", ts.expected, res)
			} else {
				t.Logf("Success!")
			}
		})
	}
}
