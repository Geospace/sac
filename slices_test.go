package sac

import "testing"

func TestEmptySlice(t *testing.T) {
	sac := New()

	k := "key"
	v := make([]string, 0)
	sac.Set(k, v)

	b := sac.GetStringSlice(k)
	trueAssert(t, b != nil)
	trueAssert(t, len(b) == 0)
	trueAssert(t, cap(b) == 0)
}

func TestSlice3(t *testing.T) {
	sac := New()

	k := "key"
	v := make([]string, 3)

	v[0] = "0"
	v[1] = "1"
	v[2] = "2"
	sac.Set(k, v)

	b := sac.GetStringSlice(k)

	trueAssert(t, b != nil)
	trueAssert(t, len(b) == 3)
	trueAssert(t, cap(b) == 3)

	simpleAssert(t, v[0], "0")
	simpleAssert(t, v[1], "1")
	simpleAssert(t, v[2], "2")
}

func TestSliceWrong(t *testing.T) {
	sac := New()

	k := "key"
	v := make([]uint8, 3)

	v[0] = 0
	v[1] = 1
	v[2] = 2
	sac.Set(k, v)

	b := sac.GetStringSlice(k)

	trueAssert(t, b != nil)
	trueAssert(t, len(b) == 0)
	trueAssert(t, cap(b) == 0)
}
