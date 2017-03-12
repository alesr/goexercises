package goexercises

import (
	"log"

	"github.com/alesr/numbers"
)

// LargestPrimeFactor - https://projecteuler.net/problem=3
// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143?
func LargestPrimeFactor(n int) int {
	factors, err := numbers.FactorsList(n)
	if err != nil {
		log.Fatal(err)
	}
	max := factors[0]
	for _, i := range factors {
		if i > max {
			max = i
		}
	}
	return max
}
