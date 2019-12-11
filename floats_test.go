package sac

import (
	"math"
	"testing"
)

func TestEmptyFloat(t *testing.T) {
	sac := New()

	k := "does_not_exist"
	v := 0

	b := sac.GetFloat(k)
	simpleAssert(t, b, float64(v))
}

func TestWrongFloat(t *testing.T) {
	sac := New()

	k := "wrong"
	v := "something_not_a_float"

	sac.Set(k, v)
	b := sac.GetFloat(k)
	simpleAssert(t, b, float64(0))
}

func TestSetGetFloat32(t *testing.T) {
	sac := New()
	k := "key"

	v := float32(math.MaxFloat32)
	sac.Set(k, v)

	f := sac.GetFloat32(k)
	simpleAssert(t, f, v)

	fAny := sac.GetFloat(k)
	// simpleAssert takes interface{} so we must cast to actual
	// expected type here
	simpleAssert(t, fAny, float64(v))
}
func TestSetGetFloat64(t *testing.T) {
	sac := New()
	k := "key"

	v := float64(math.MaxFloat64)
	sac.Set(k, v)

	f := sac.GetFloat64(k)
	simpleAssert(t, f, v)

	fAny := sac.GetFloat(k)
	simpleAssert(t, fAny, v)
}
