package sac

import "testing"

func TestDeepString(t *testing.T) {
	sac := New()

	k := "some.nested.string"
	v := "value"

	sac.Set(k, v)
	str := sac.GetString(k)
	simpleAssert(t, str, v)
}

func TestDeepNumber(t *testing.T) {
	sac := New()

	k := "some.nested.number"
	v := 3

	sac.Set(k, v)
	str := sac.GetNumber(k)
	simpleAssert(t, str, v)
}

func TestDeepFloat(t *testing.T) {
	sac := New()

	k := "some.nested.float"
	v := 3.14

	sac.Set(k, v)
	str := sac.GetFloat(k)
	simpleAssert(t, str, v)
}

func TestDeepBoolean(t *testing.T) {
	sac := New()

	k := "some.nested.boolean"
	v := true

	sac.Set(k, v)
	str := sac.GetBool(k)
	simpleAssert(t, str, v)
}
