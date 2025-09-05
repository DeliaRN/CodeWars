package kata12_greedisgood

import "testing"

func TestScore(t *testing.T) {

	tests := []struct {
		dices    [5]int
		expected int
	}{
		{[5]int{4, 4, 4, 3, 3}, 400},
		{[5]int{1, 1, 1, 1, 1}, 1200},
		{[5]int{2, 4, 4, 5, 4}, 450},
		{[5]int{1, 1, 1, 5, 1}, 1150},
		{[5]int{2, 3, 4, 6, 2}, 0},
		{[5]int{0, 5, 2, 1, 0}, -1},
	}

	for _, test := range tests {
		result := Score(test.dices)
		if result != test.expected {
			t.Errorf("got %d want %d", result, test.expected)
		}
	}

}
