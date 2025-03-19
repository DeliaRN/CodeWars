package kata02_repeatString

import "testing"

func TestRepeatString(t *testing.T) {
	numberOfTimes := 5
	str := "a"

	got := RepeatString(numberOfTimes, str)
	want := "aaaaa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
