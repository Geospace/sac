package sac

import "testing"

func TestEmptyBool(t *testing.T) {
	sac := New()

	k := "does_not_exist"
	v := false

	b := sac.GetBool(k)
	simpleAssert(t, b, v)
}

func TestWrongBool(t *testing.T) {
	sac := New()

	k := "wrong"
	v := "something_not_a_bool"

	sac.Set(k, v)
	b := sac.GetBool(k)
	simpleAssert(t, b, false)
}

func TestSetGetFlatBool(t *testing.T) {
	sac := New()

	k := "key"
	v := true

	sac.Set(k, v)
	b := sac.GetBool(k)
	simpleAssert(t, b, v)
}
