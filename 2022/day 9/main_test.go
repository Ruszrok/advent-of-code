package main

import (
	"fmt"
	"testing"
)

func TestIsValidRope(t *testing.T) {
	tests := []struct {
		head     Coords
		tail     Coords
		expected bool
	}{
		{[2]int{0, 0}, [2]int{0, 0}, true},
		{[2]int{1, 0}, [2]int{0, 0}, true},
		{[2]int{-1, 0}, [2]int{0, 0}, true},
		{[2]int{0, 1}, [2]int{0, 0}, true},
		{[2]int{0, -1}, [2]int{0, 0}, true},
		{[2]int{1, 1}, [2]int{0, 0}, true},
		{[2]int{1, -1}, [2]int{0, 0}, true},
		{[2]int{-1, 1}, [2]int{0, 0}, true},
		{[2]int{-1, -1}, [2]int{0, 0}, true},
		{[2]int{0, 2}, [2]int{0, 0}, false},
		{[2]int{2, 0}, [2]int{0, 0}, false},
		{[2]int{-2, 0}, [2]int{0, 0}, false},
		{[2]int{0, -2}, [2]int{0, 0}, false},
		{[2]int{3, 3}, [2]int{0, 0}, false},
		{[2]int{2, 1}, [2]int{0, 0}, false},
		{[2]int{1, 2}, [2]int{0, 0}, false},
	}

	for _, ts := range tests {
		t.Run(fmt.Sprintf("Get code for input: %v, %v", ts.head, ts.tail), func(t *testing.T) {
			r := InitRope()
			r.head = ts.head
			r.tail = ts.tail
			res := r.IsValid()
			if res != ts.expected {
				t.Fatalf("Wrong anser. Expected = %t. Got = %t", ts.expected, res)
			} else {
				t.Logf("Success!")
			}
		})
	}
}
