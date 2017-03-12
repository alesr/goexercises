package goexercises

import "testing"

func TestEvenFibonacci(t *testing.T) {
	testCases := []struct {
		input, expected int
	}{
		{3, 2},
		{10, 10},
		{4000000, 4613732},
	}
	for _, test := range testCases {
		observed := EvenFibonacci(test.input)
		if observed != test.expected {
			t.Errorf("for input '%d', expected '%d', observed '%d'\n",
				test.input, test.expected, observed)
		}
	}
}

func BenchmackEvenFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EvenFibonacci(10000)
	}
}
