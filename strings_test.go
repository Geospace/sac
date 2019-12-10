package sac

import "testing"

func TestSetGetString(t *testing.T) {
	sac := New()

	k := "key"
	v := "value"

	sac.Set(k, v)
	str := sac.GetString(k)
	simpleAssert(t, str, v)
}
