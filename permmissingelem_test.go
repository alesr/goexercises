package goexercises

import "testing"

func TestPermMissingElem(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{2, 3, 1, 5}, 4},
		{[]int{2, 1, 4}, 3},
		{[]int{2, 3, 1}, 4},
	}
	for _, test := range testCases {
		observed := PermMissingElem(test.input)
		if observed != test.expected {
			t.Errorf("input '%v', expected '%d', observed '%d'\n",
				test.input, test.expected, observed)
		}
	}
}

func BenchmarkPermMissingElem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PermMissingElem([]int{2, 3, 1, 5})
	}
}
