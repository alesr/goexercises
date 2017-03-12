package goexercises

import "math"

// FrogJmp - https://codility.com/programmers/lessons/3-time_complexity/frog_jmp/
// Count minimal number of jumps from position X to Y.
func FrogJmp(x, y, d int) int {
	return int(math.Ceil((float64(y - x)) / float64(d)))
}
