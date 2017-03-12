package goexercises

import (
	"testing"
)

func TestSmallestMultiple(t *testing.T) {
	testCases := []struct {
		input, expected int
	}{
		{10, 2520},
		{20, 232792560},
	}
	for _, test := range testCases {
		observed := SmallestMultiple(test.input)
		if observed != test.expected {
			t.Errorf("for input '%d', expected '%d', observed: '%d'\n",
				test.input, test.expected, observed)
		}
	}
}

func BenchmarkSmallestMultiple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SmallestMultiple(20)
	}
}
