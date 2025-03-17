package kata4

import (
	"reflect"
	"testing"
)

func TestDescArray(t *testing.T) {
	got := descendingArray(5)
	want := []int{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
