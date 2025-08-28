package kata12_yesorno

import "testing"

func TestIsItTrue(t *testing.T) {

	t.Run("is it true?", func(t *testing.T) {
		got := isItTrue(true)
		want := "Yes"
		checkTruth(t, got, want)
	})

	t.Run("maybe false", func(t *testing.T) {
		got := isItTrue(false)
		want := "No"
		checkTruth(t, got, want)
	})

}

func checkTruth(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
