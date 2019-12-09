package sac

import (
	"testing"
)

func simpleAssert(t *testing.T, got interface{}, want interface{}) {
	if want != got {
		t.Fatalf("got: %v, want: %v", got, want)
	}
}

func trueAssert(t *testing.T, got bool) {
	want := true

	if got != want {
		t.Fatalf("got: %v, want: %v", got, want)
	}
}

func errAssert(t *testing.T, err error) {
	if err != nil {
		t.Fatalf(err.Error())
	}
}
