package sac

import "testing"

func TestNew(t *testing.T) {
	sac := New()
	trueAssert(t, sac != nil)
}

func TestSetGetString(t *testing.T) {
	sac := New()

	k := "key"
	v := "value"

	sac.Set(k, v)
	str := sac.GetString(k)
	simpleAssert(t, str, v)

	empty := sac.GetString("doesNotExist")
	simpleAssert(t, empty, "")

	sac.Set(k, 1337)
	wrong := sac.GetString(k)
	simpleAssert(t, wrong, "")
}

func TestSetGetStringDeep(t *testing.T) {
	sac := New()

	k := "some.nested.key"
	v := "value"

	sac.Set(k, v)
	str := sac.GetString(k)
	simpleAssert(t, str, v)

	empty := sac.GetString("does.not.exist")
	simpleAssert(t, empty, "")

	sac.Set(k, 1337)
	wrong := sac.GetString(k)
	simpleAssert(t, wrong, "")
}

func TestSetGetNumber(t *testing.T) {
	sac := New()

	k := "key"
	v := 3

	sac.Set(k, v)
	str := sac.GetNumber(k)
	simpleAssert(t, str, v)

	empty := sac.GetNumber("doesNotExist")
	simpleAssert(t, empty, 0)

	sac.Set(k, "this is wrong")
	wrong := sac.GetNumber(k)
	simpleAssert(t, wrong, 0)
}

func TestSetGetNumberDeep(t *testing.T) {
	sac := New()

	k := "some.nested.key"
	v := 3

	sac.Set(k, v)
	str := sac.GetNumber(k)
	simpleAssert(t, str, v)

	empty := sac.GetNumber("does.not.exist")
	simpleAssert(t, empty, 0)

	sac.Set(k, "this is wrong")
	wrong := sac.GetNumber(k)
	simpleAssert(t, wrong, 0)
}

func TestSetGetBool(t *testing.T) {
	sac := New()

	k := "key"
	v := true

	sac.Set(k, v)
	str := sac.GetBool(k)
	simpleAssert(t, str, v)

	empty := sac.GetBool("doesNotExist")
	simpleAssert(t, empty, false)

	sac.Set(k, "this is wrong")
	wrong := sac.GetBool(k)
	simpleAssert(t, wrong, false)
}

func TestSetGetBoolDeep(t *testing.T) {
	sac := New()

	k := "some.nested.key"
	v := true

	sac.Set(k, v)
	str := sac.GetBool(k)
	simpleAssert(t, str, v)

	empty := sac.GetBool("does.not.exist")
	simpleAssert(t, empty, false)

	sac.Set(k, "this is wrong")
	wrong := sac.GetBool(k)
	simpleAssert(t, wrong, false)
}
