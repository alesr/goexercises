package goexercises

// PermMissingElem - https://codility.com/programmers/lessons/3-time_complexity/perm_missing_elem/
// Find the missing element in a given permutation.
func PermMissingElem(arr []int) int {
	// Wikipedia: infinite series
	len := len(arr) + 1
	total := len * (len + 1) / 2
	for _, v := range arr {
		total -= v
	}
	return total
}
