package goexercises

import "strconv"

// FizzBuzzWoof receives n, d1, d2, d3 as positive integer as arguments and returns:
//	Fizz -> if n is divisible by d1
//	Buzz -> if n is divisible by d2
//	Woof -> if n is divisible by d3
//	FizzBuzz -> if n is divisible by d1 and d2
//	FizzWoof -> if n is divisible by d1 and d3
//	BuzzWoof -> if n is divisible by d2 and d3
//	FizzBuzzWoof -> if n is divisible by d1 and d2 and d3
//	string(n) -> if n is not evenly divisible by any given divisor
func FizzBuzzWoof(n, d1, d2, d3 int) string {
	switch {
	case n%(d1*d2*d3) == 0:
		return "FizzBuzzWoof"
	case n%(d2*d3) == 0:
		return "BuzzWoof"
	case n%(d1*d3) == 0:
		return "FizzWoof"
	case n%(d1*d2) == 0:
		return "FizzBuzz"
	case n%(d3) == 0:
		return "Woof"
	case n%(d2) == 0:
		return "Buzz"
	case n%(d1) == 0:
		return "Fizz"
	default:
		return strconv.Itoa(n)
	}
}
