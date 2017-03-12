package goexercises

import "testing"

func TestTapeEquilibrium(t *testing.T) {
	testCases := []struct {
		slice    []int
		expected int
	}{
		{[]int{3, 1, 2, 4, 3}, 1},
		{[]int{3, 1, 2, 4, 3, 1, 2, 3, 10}, 1},
		{[]int{2, 3, 10}, 5},
		{[]int{0}, 0},
		{[]int{0, 1}, 1},
	}
	for _, test := range testCases {
		observed := TapeEquilibrium(test.slice)
		if observed != test.expected {
			t.Errorf("for slice '%v', expected '%d', observed '%d'\n",
				test.slice, test.expected, observed)
		}
	}
}

func BenchmarkTapeEquilibrium(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TapeEquilibrium([]int{0, 1, 2, 3, 4, 5})
	}
}

func TestSumSliceInt(t *testing.T) {
	testCases := []struct {
		slice    []int
		expected int
	}{
		{[]int{1, 1, 1}, 3},
		{[]int{0}, 0},
		{[]int{1}, 1},
		{[]int{1, 3}, 4},
		{[]int{1, -1}, 0},
		{[]int{-1, -1}, -2},
	}
	for _, test := range testCases {
		observed := sumSliceInt(test.slice)
		if observed != test.expected {
			t.Errorf("for slice '%v', expected '%d', observed '%d'\n",
				test.slice, test.expected, observed)
		}
	}
}

func BenchmarkSumSliceInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumSliceInt([]int{0, 1, 2, 3, 4, 5})
	}
}
