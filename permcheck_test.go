package goexercises

import "testing"

func TestPermCheck(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{4, 1, 3, 2}, 1},
	}

	for _, test := range testCases {
		observed := PermCheck(test.input)
		if observed != test.expected {
			t.Errorf("input '%v', expected '%d', observed '%d'\n",
				test.input, test.expected, observed)
		}
	}
}

func BenchmarkPermCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PermCheck([]int{4, 1, 3, 2})
	}
}
