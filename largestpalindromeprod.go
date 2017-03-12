package goexercises

import (
	"strconv"

	"github.com/alesr/strings"
)

// LargestPalindromeProd - https://projecteuler.net/problem=4
// A palindromic number reads the same both ways.
// The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.
func LargestPalindromeProd(n int) int {
	palindrome, floor := 0, n/10
	// Starting with n x n
	for i := n; i > floor; i-- {
		for j := n; j > floor; j-- {
			product := i * j
			strProd := strconv.Itoa(product) // convert int into string.
			revStr := strings.Flip(strProd)
			// Check equality and certifies the is the greater product.
			if strProd == revStr && product > palindrome {
				palindrome = product
				break
			}
		}
	}
	return palindrome
}
