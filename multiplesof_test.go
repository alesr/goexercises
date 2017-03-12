package goexercises

import "testing"

func TestMultiplesOf(t *testing.T) {
	testCases := []struct {
		input, expected int
	}{
		{10, 23},
		{1000, 233168},
	}
	for _, test := range testCases {
		observed := MultiplesOf(test.input)
		if observed != test.expected {
			t.Errorf("input '%d', expected '%d', observed '%d'\n",
				test.input, test.expected, observed)
		}
	}
}

func BenchmarkMultiplesOf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MultiplesOf(1000)
	}
}
