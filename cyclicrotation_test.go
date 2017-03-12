package goexercises

import (
	"reflect"
	"testing"
)

func TestCyclicRotation(t *testing.T) {
	testCases := []struct {
		input    []int
		rotation int
		expected []int
	}{
		{[]int{3, 8, 9, 7, 6}, 0, []int{3, 8, 9, 7, 6}},
		{[]int{3, 8, 9, 7, 6}, 1, []int{6, 3, 8, 9, 7}},
		{[]int{3, 8, 9, 7, 6}, 3, []int{9, 7, 6, 3, 8}},
	}
	for _, test := range testCases {
		observed := CyclicRotation(test.input, test.rotation)
		if !reflect.DeepEqual(observed, test.expected) {
			t.Errorf("for input '%v', rotation '%d', expected '%v', observed '%v'\n",
				test.input, test.rotation, test.expected, observed)
		}
	}
}

func BenchmarkCyclicRotation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CyclicRotation([]int{3, 8, 9, 7, 6}, 3)
	}
}
