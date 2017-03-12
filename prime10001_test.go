package goexercises

import "testing"

func TestPrime10001(t *testing.T) {
	testCases := []struct {
		input, expected int
	}{
		{6, 13},
		{10001, 104743},
	}
	for _, test := range testCases {
		observed := Prime10001(test.input)
		if observed != test.expected {
			t.Errorf("for input '%d', expected '%d', observed '%d'\n",
				test.input, test.expected, observed)
		}
	}
}

func BenchmarkPrime10001(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Prime10001(10001)
	}
}
