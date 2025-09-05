package kata13sumintervals

import "testing"

func TestSumOfIntervals(t *testing.T) {

	tests := []struct {
		intervals [][2]int
		expected  int
	}{
		{[][2]int{{1, 5}}, 4},
		{[][2]int{{1, 5}, {6, 10}}, 8},
		{[][2]int{{1, 5}, {1, 5}}, 4},
	}

	for _, test := range tests {
		result := SumOfIntervals(test.intervals)
		if result != test.expected {
			t.Errorf("got %d want %d", result, test.expected)
		}
	}
}
