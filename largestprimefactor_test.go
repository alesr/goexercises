package goexercises

import "testing"

func TestPrimeFactor(t *testing.T) {
	testCases := []struct {
		input, expected int
	}{
		{13195, 29},
		{600851475143, 6857},
	}
	for _, test := range testCases {
		observed := LargestPrimeFactor(test.input)
		if observed != test.expected {
			t.Errorf("for input '%d', expected '%d', observed '%d'\n",
				test.input, test.expected, observed)
		}
	}
}

func BenchmarkLargestPrimeFactor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LargestPrimeFactor(600851475143)
	}
}
