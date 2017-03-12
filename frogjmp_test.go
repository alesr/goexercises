package goexercises

import "testing"

func TestFrogJmp(t *testing.T) {
	testCases := []struct {
		x, y, d, expected int
	}{
		{10, 85, 30, 3},
	}
	for _, test := range testCases {
		observed := FrogJmp(test.x, test.y, test.d)
		if observed != test.expected {
			t.Errorf("x '%d', y '%d', d '%d', expected '%d', observed '%d'\n",
				test.x, test.y, test.d, test.expected, observed)
		}
	}
}

func BenchmarkFrogJmp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FrogJmp(10, 85, 30)
	}
}
