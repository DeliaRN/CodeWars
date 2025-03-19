package kata9_stringtonumb

import (
	"reflect"
	"testing"
)

func TestConversor(t *testing.T) {

	got := stringConversor("1234")
	want := 1234

	if !reflect.DeepEqual(got, want) {
		t.Errorf(" want %v got %q", got, want)
	}

}
