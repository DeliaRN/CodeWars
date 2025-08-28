package kata13_oddEven

import "testing"

func TestOddEven(t *testing.T) {
	got := oddOrEven(8)
	want := "Even"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
