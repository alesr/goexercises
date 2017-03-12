package goexercises

// OddOccurrences - https://codility.com/programmers/lessons/2-arrays/odd_occurrences_in_array/start/
// Find value that occurs in odd number of elements.
func OddOccurrences(arr []int) int {
	m := map[int]int{}
	for _, v := range arr {
		m[v]++
	}
	for k, v := range m {
		if v%2 != 0 {
			return k
		}
	}
	return 0
}
