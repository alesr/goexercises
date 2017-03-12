package goexercises

import "testing"

func TestFizzBuzzWoof(t *testing.T) {
	testCases := []struct {
		n, d1, d2, d3 int
		expected      string
	}{
		{1, 3, 5, 7, "1"},
	}
	for _, test := range testCases {
		observed := FizzBuzzWoof(test.n, test.d1, test.d2, test.d3)
		if observed != test.expected {
			t.Errorf("for n '%d', d1 '%d', d2 '%d', d3 '%d, expected '%s', observed '%s'\n",
				test.n, test.d1, test.d2, test.d3, test.expected, observed)
		}
	}
}

func BenchmarkFizzBuzzWoof(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FizzBuzzWoof(1000, 3, 5, 7)
	}
}
