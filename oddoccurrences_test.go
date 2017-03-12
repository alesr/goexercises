package goexercises

import "testing"

func TestOddOccurrences(t *testing.T) {
	testCases := []struct {
		arr      []int
		expected int
	}{
		{[]int{9, 3, 9, 3, 9, 7, 9}, 7},
	}
	for _, test := range testCases {
		observed := OddOccurrences(test.arr)
		if observed != test.expected {
			t.Errorf("for list '%v', expected '%d', observed '%d'\n",
				test.arr, test.expected, observed)
		}
	}
}

func BenchmarkOddOccurrences(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OddOccurrences([]int{9, 3, 9, 3, 9, 7, 9})
	}
}
