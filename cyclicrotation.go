package goexercises

// CyclicRotation - https://codility.com/programmers/lessons/2-arrays/cyclic_rotation/
// Rotate an array to the right by a given number of steps.
func CyclicRotation(arr []int, rotation int) []int {
	if len(arr) <= 1 || rotation == 0 {
		return arr
	}
	// a is the last element followed by the n elements
	// previously before the last element
	for i := 0; i < rotation; i++ {
		arr = append(arr[len(arr)-1:], arr[:len(arr)-1]...)
	}
	return arr
}
