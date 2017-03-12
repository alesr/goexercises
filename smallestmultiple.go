package goexercises

// SmallestMultiple - https://projecteuler.net/problem=5
// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
func SmallestMultiple(n int) int {
	flag := false
	r, k := 0, 1
	for !flag {
		i := 1
		// Iterates over all the divisor we need to try
		for i <= n {
			// If divisible, keep incrementing the i
			if k%i == 0 {
				flag = true
				// We hit the target
				if i == n {
					r = k
					goto end
				}
			} else {
				flag = false
				break
			}
			i++
		}
		k++
	}
end:
	return r
}
