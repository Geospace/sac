package sac

import "testing"

func simpleAssert(t *testing.T, got interface{}, want interface{}) {
	if want != got {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func trueAssert(t *testing.T, got bool) {
	want := true

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
