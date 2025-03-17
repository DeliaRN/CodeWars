package kata5

import "testing"

func TestHopper(t *testing.T) {
	got := hopper(8)
	want := 36

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

}
