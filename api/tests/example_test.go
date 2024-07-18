package tests

import (
	"testing"
)

func TestExample(t *testing.T) {
	got := 4
	want := 4

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	} else {
		t.Logf("got %q, wanted %q", got, want)
	}
}
