package goexercises

import "testing"

func TestBinaryGap(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{5, 1},
		{1, 0},
		{4, 0},
		{10, 1},
		{9, 2},
	}
	for _, test := range testCases {
		observed := BinaryGap(test.input)
		if observed != test.expected {
			t.Errorf("for input '%d', expected '%d', observed '%d'\n",
				test.input, test.expected, observed)
		}
	}
}

func BenchmarkBinaryGap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinaryGap(10000)
	}
}

func TestCountGaps(t *testing.T) {
	var testCases = []struct {
		input    string
		expected int
	}{
		{"10001", 3},
		{"1", 0},
		{"10", 1},
		{"101", 1},
		{"0", 1},
	}
	for _, test := range testCases {
		observed := countGaps(test.input)
		if observed != test.expected {
			t.Errorf("for input '%s', expected '%d', observed '%d'\n",
				test.input, test.expected, observed)
		}
	}
}

func BenchmarkCountGaps(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countGaps("1000001111011110110001")
	}
}
