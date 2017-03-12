package goexercises

import (
	"github.com/alesr/numbers"
)

// Prime10001 - https://projecteuler.net/problem=7
// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
// we can see that the 6th prime is 13.
// What is the 10 001st prime number?
func Prime10001(n int) int {
	var count int
	// We start at 3, so we're one iteration ahead.
	for i := 3; ; i += 2 {
		if numbers.IsPrime(i) {
			count++
			if count == n-1 { // minus one iteration
				return i
			}
		}
	}
}
