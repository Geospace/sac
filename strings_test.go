package sac

import "testing"

func TestEmptyString(t *testing.T) {
	sac := New()

	k := "does_not_exist"
	v := ""

	b := sac.GetString(k)
	simpleAssert(t, b, v)
}

func TestWrongString(t *testing.T) {
	sac := New()

	k := "wrong"
	v := 1337

	sac.Set(k, v)
	b := sac.GetString(k)
	simpleAssert(t, b, "")
}
func TestSetGetString(t *testing.T) {
	sac := New()

	k := "key"
	v := "value"

	sac.Set(k, v)
	str := sac.GetString(k)
	simpleAssert(t, str, v)
}
