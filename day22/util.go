// Utility functions

package main

import (
	"strconv"
)

// Parse an integer
func atoi(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		panic("Bad number")
	}
	return i
}

// Simple integer min/max
func min(a, b int64) int64 {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}

//  Min/max of a list of numbers
func minmax(l []int64) (int64, int64) {
	var minVal, maxVal int64
	for i, n := range l {
		if i == 0 || n < minVal {
			minVal = n
		}
		if i == 0 || n > maxVal {
			maxVal = n
		}
	}
	return minVal, maxVal
}
