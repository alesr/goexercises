package goexercises

import (
	"testing"
)

func TestSumSqrDiff(t *testing.T) {
	testCases := []struct {
		input, expected int
	}{
		{10, 2640},
		{100, 25164150},
	}
	for _, test := range testCases {
		observed := SumSqrDiff(test.input)
		if observed != test.expected {
			t.Errorf("for input '%d', expected '%d, observed '%d'\n",
				test.input, test.expected, observed)
		}
	}
}

func BenchmarkSumSqrDiff(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumSqrDiff(100)
	}
}
