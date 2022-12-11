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

func TestIsValidOneTimeMoveInMultiKnotRope(t *testing.T) {
	tests := []struct {
		knots     []Coords
		direction string
		expected  []Coords
	}{
		{[]Coords{{0, 0}, {0, 0}, {0, 0}}, "R", []Coords{{1, 0}, {0, 0}, {0, 0}}},
		{[]Coords{{1, 0}, {0, 0}, {0, 0}}, "R", []Coords{{2, 0}, {1, 0}, {0, 0}}},
		{[]Coords{{2, 0}, {1, 0}, {0, 0}}, "R", []Coords{{3, 0}, {2, 0}, {1, 0}}},
		{[]Coords{{2, 0}, {1, 0}, {0, 0}}, "L", []Coords{{1, 0}, {1, 0}, {0, 0}}},
		{[]Coords{{1, 0}, {1, 0}, {0, 0}}, "L", []Coords{{0, 0}, {1, 0}, {0, 0}}},
		{[]Coords{{0, 0}, {0, 0}, {0, 0}}, "U", []Coords{{0, 1}, {0, 0}, {0, 0}}},
		{[]Coords{{0, 1}, {0, 0}, {0, 0}}, "U", []Coords{{0, 2}, {0, 1}, {0, 0}}},
		{[]Coords{{0, 2}, {0, 1}, {0, 0}}, "U", []Coords{{0, 3}, {0, 2}, {0, 1}}},
		{[]Coords{{3, 0}, {2, 0}, {1, 0}, {0, 0}, {0, 0}}, "U", []Coords{{3, 1}, {2, 0}, {1, 0}, {0, 0}, {0, 0}}},
		{[]Coords{{3, 1}, {2, 0}, {1, 0}, {0, 0}, {0, 0}}, "U", []Coords{{3, 2}, {3, 1}, {2, 1}, {1, 1}, {0, 0}}},
		{[]Coords{{3, 2}, {3, 1}, {2, 1}, {1, 1}, {0, 0}}, "U", []Coords{{3, 3}, {3, 2}, {2, 1}, {1, 1}, {0, 0}}},
		{[]Coords{{1, 2}, {2, 1}, {2, 0}, {1, 0}}, "L", []Coords{{0, 2}, {1, 2}, {1, 1}, {1, 0}}},
	}

	for _, ts := range tests {
		t.Run(fmt.Sprintf("Get code for input: %v direction: %s", ts.knots, ts.direction), func(t *testing.T) {
			r := InitMultiKnotRope(len(ts.knots))
			r.knots = ts.knots
			r.oneTimeMove(ts.direction)
			if !isEqual(r.knots, ts.expected) {
				t.Fatalf("Wrong anser. Expected = %v. Got = %v", ts.expected, r.knots)
			} else {
				t.Logf("Success!")
			}
		})
	}
}

func isEqual(left []Coords, right []Coords) bool {
	if len(left) != len(right) {
		return false
	}

	for i := 0; i < len(left); i++ {
		if left[i][0] != right[i][0] || left[i][1] != right[i][1] {
			return false
		}
	}

	return true
}
