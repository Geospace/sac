package sac

import (
	"fmt"
	"runtime/debug"
	"testing"
)

func testFailedTrace(t *testing.T, str string) {
	debug.PrintStack()
	t.Fatalf(str)
}

func simpleAssert(t *testing.T, got interface{}, want interface{}) {
	if want != got {
		testFailedTrace(t, fmt.Sprintf("got: %v, want: %v", got, want))
	}
}

func trueAssert(t *testing.T, got bool) {
	want := true

	if got != want {
		debug.PrintStack()
		testFailedTrace(t, fmt.Sprintf("got: %v, want: %v", got, want))
	}
}

func errAssert(t *testing.T, err error) {
	if err != nil {
		testFailedTrace(t, err.Error())
	}
}
