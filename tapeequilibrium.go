package goexercises

import "math"

// TapeEquilibrium - https://codility.com/programmers/lessons/3-time_complexity/tape_equilibrium/
// Minimize the value `|(A[0] + ... + A[P-1]) - (A[P] + ... + A[N-1])|`
func TapeEquilibrium(slice []int) int {
	total := sumSliceInt(slice)
	min, leftSide := math.MaxFloat64, 0
	for _, v := range slice {
		rightSide := total - v - leftSide
		leftSide += v
		difference := math.Abs(float64(rightSide - leftSide))
		if difference < min {
			min = difference
		}
	}
	return int(min)
}

func sumSliceInt(slice []int) (sum int) {
	for _, v := range slice {
		sum += v
	}
	return
}
