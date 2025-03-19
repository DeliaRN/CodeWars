package kata11_initialscreator

import "testing"

func TestInitialsCreator(t *testing.T) {

	got := initialsCreator("Sam Harris")
	want := "S.H"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
