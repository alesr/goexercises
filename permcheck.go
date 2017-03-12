package goexercises

// PermCheck - https://codility.com/programmers/lessons/4-counting_elements/perm_check/
// Check whether array A is a permutation.
func PermCheck(slice []int) int {
	len := len(slice)
	counter := make([]int, len)
	for _, v := range slice {
		if v < 1 || v > len {
			return 0
		}
		if counter[v-1] == 1 {
			return 0
		}
		counter[v-1] = 1
	}
	return 1
}
