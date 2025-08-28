package kata14_stringToArray

import (
	"reflect"
	"testing"
)

func TestStringToArray(t *testing.T) {
	got := StringToArray("Hola caracola")
	want := []string{"Hola", "caracola"}

	if reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
