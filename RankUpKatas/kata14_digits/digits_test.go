package kata14digits

import "testing"

func TestDigits(t *testing.T) {

	tests := []struct {
		num      int
		expected []int
	}{
		{12345, []int{3, 4, 5, 6, 5, 6, 7, 7, 8, 9}},
		{0, []int{0}},
	}

	for _, test := range tests {
		result := Digits(test.num)
		if !equal(result, test.expected) {
			t.Errorf("got %v want %v", result, test.expected)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
