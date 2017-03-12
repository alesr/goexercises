package goexercises

import (
	"testing"
)

func TestLargestPalindromeProd(t *testing.T) {
	testCases := []struct {
		input, expected int
	}{
		{99, 9009},
		{999, 906609},
	}
	for _, test := range testCases {
		observed := LargestPalindromeProd(test.input)
		if observed != test.expected {
			t.Errorf("for input '%d', expected '%d', observed '%d'\n",
				test.input, test.expected, observed)
		}
	}
}

func BenchmarkLargestPalindromeProd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LargestPalindromeProd(999)
	}
}
