package goexercises

import (
	"strconv"
	"strings"
)

// BinaryGap - https://codility.com/programmers/lessons/1-iterations/binary_gap/
// Find longest sequence of zeros in binary representation of an integer.
func BinaryGap(n int) int {
	// Binary representation of n
	binary := strconv.FormatInt(int64(n), 2)
	// Remove trailing zeros
	binary = binary[:strings.LastIndex(binary, "1")+1]
	// Check if binary contains zeros
	if !strings.ContainsAny(binary, "0") {
		return 0
	}
	return countGaps(binary)
}

func countGaps(s string) int {
	var max, counter int
	for _, char := range s {
		if char == '0' {
			counter++
			if counter > max {
				max = counter
			}
		} else {
			counter = 0
		}
	}
	return max
}
