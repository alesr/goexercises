package goexercises

// SumSqrDiff - https://projecteuler.net/problem=6
// The sum of the squares of the first ten natural numbers is,
// 12 + 22 + ... + 102 = 385
// The square of the sum of the first ten natural numbers is,
// (1 + 2 + ... + 10)2 = 552 = 3025
// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
func SumSqrDiff(n int) int {
	sumOfSq := 0 // Sum of the squares ie. 1^2+2^3...
	sqSum := 0   // Square of the sum ie.  (1+2+...)^3
	for i := 1; i <= n; i++ {
		sumOfSq += i * i
		sqSum += i
	}
	return sqSum*sqSum - sumOfSq
}
