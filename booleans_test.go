package sac

import "testing"

func TestSetGetFlatBool(t *testing.T) {
	sac := New()

	k := "key"
	v := true

	sac.Set(k, v)
	b := sac.GetBool(k)
	simpleAssert(t, b, v)
}
