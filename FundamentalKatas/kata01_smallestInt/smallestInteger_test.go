package kata01_smallestInt

import "testing"

func TestSmallestIntegerFinder(t *testing.T) {

	tests := []struct {
		numbers  []int
		expected int
	}{
		{[]int{55, 96, 66, 1}, 1},
		{[]int{34, -345, -1, -100}, -345},
	}

	for _, test := range tests {
		result := SmallestIntegerFinder(test.numbers)
		if result != test.expected {
			t.Errorf("got %d want %d", result, test.expected)
		}
	}

}
