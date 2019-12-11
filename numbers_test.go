package sac

import (
	"math"
	"testing"
)

// TODO: Could add more tests for the 0 values returns, but let's wait
// for the safe functions equivalents

func TestEmptyNumber(t *testing.T) {
	sac := New()

	k := "does_not_exist"
	v := 0

	b := sac.GetNumber(k)
	simpleAssert(t, b, v)
}

func TestWrongNumber(t *testing.T) {
	sac := New()

	k := "wrong"
	v := "something_not_a_number"

	sac.Set(k, v)
	b := sac.GetNumber(k)
	simpleAssert(t, b, 0)
}

func TestSetGetNumberU8(t *testing.T) {
	sac := New()
	k := "key"

	v := uint8(math.MaxUint8)
	sac.Set(k, v)

	n := sac.GetUint8(k)
	simpleAssert(t, n, v)

	nAny := sac.GetNumber(k)
	simpleAssert(t, nAny, int(v))
}

func TestSetGetNumber8Max(t *testing.T) {
	sac := New()
	k := "key"

	v := int8(math.MaxInt8)
	sac.Set(k, v)

	n := sac.GetInt8(k)
	simpleAssert(t, n, v)

	nAny := sac.GetNumber(k)
	simpleAssert(t, nAny, int(v))
}

func TestSetGetNumber8Min(t *testing.T) {
	sac := New()
	k := "key"

	v := int8(math.MinInt8)
	sac.Set(k, v)

	n := sac.GetInt8(k)
	simpleAssert(t, n, v)

	nAny := sac.GetNumber(k)
	simpleAssert(t, nAny, int(v))
}

func TestSetGetNumberU16(t *testing.T) {
	sac := New()
	k := "key"

	v := uint16(math.MaxUint16)
	sac.Set(k, v)

	n := sac.GetUint16(k)
	simpleAssert(t, n, v)

	nAny := sac.GetNumber(k)
	simpleAssert(t, nAny, int(v))
}

func TestSetGetNumber16Max(t *testing.T) {
	sac := New()
	k := "key"

	v := int16(math.MaxInt16)
	sac.Set(k, v)

	n := sac.GetInt16(k)
	simpleAssert(t, n, v)

	nAny := sac.GetNumber(k)
	simpleAssert(t, nAny, int(v))
}

func TestSetGetNumber16Min(t *testing.T) {
	sac := New()
	k := "key"

	v := int16(math.MinInt16)
	sac.Set(k, v)

	n := sac.GetInt16(k)
	simpleAssert(t, n, v)

	nAny := sac.GetNumber(k)
	simpleAssert(t, nAny, int(v))
}

func TestSetGetNumberU32(t *testing.T) {
	sac := New()
	k := "key"

	v := uint32(math.MaxUint32)
	sac.Set(k, v)

	n := sac.GetUint32(k)
	simpleAssert(t, n, v)

	nAny := sac.GetNumber(k)
	simpleAssert(t, nAny, int(v))
}

func TestSetGetNumberU32_(t *testing.T) {
	sac := New()
	k := "key"

	v := uint(math.MaxUint32)
	sac.Set(k, v)

	nAny := sac.GetNumber(k)
	simpleAssert(t, nAny, int(v))
}

func TestSetGetNumber32Max(t *testing.T) {
	sac := New()
	k := "key"

	v := int32(math.MaxInt32)
	sac.Set(k, v)

	n := sac.GetInt32(k)
	simpleAssert(t, n, v)

	nAny := sac.GetNumber(k)
	simpleAssert(t, nAny, int(v))
}

func TestSetGetNumber32Max_(t *testing.T) {
	sac := New()
	k := "key"

	v := int(math.MaxInt32)
	sac.Set(k, v)

	nAny := sac.GetNumber(k)
	simpleAssert(t, nAny, v)
}

func TestSetGetNumber32Min(t *testing.T) {
	sac := New()
	k := "key"

	v := int32(math.MinInt32)
	sac.Set(k, v)

	n := sac.GetInt32(k)
	simpleAssert(t, n, v)

	nAny := sac.GetNumber(k)
	simpleAssert(t, nAny, int(v))
}

func TestSetGetNumber32Min_(t *testing.T) {
	sac := New()
	k := "key"

	v := int(math.MinInt32)
	sac.Set(k, v)

	nAny := sac.GetNumber(k)
	simpleAssert(t, nAny, v)
}

func TestSetGetNumberU64(t *testing.T) {
	sac := New()
	k := "key"

	v := uint64(math.MaxUint64)
	sac.Set(k, v)

	n := sac.GetUint64(k)
	simpleAssert(t, n, v)

	nAny := sac.GetNumber(k)
	simpleAssert(t, nAny, int(v))
}

func TestSetGetNumber64Max(t *testing.T) {
	sac := New()
	k := "key"

	v := int64(math.MaxInt64)
	sac.Set(k, v)

	n := sac.GetInt64(k)
	simpleAssert(t, n, v)

	nAny := sac.GetNumber(k)
	simpleAssert(t, nAny, int(v))
}

func TestSetGetNumber64Min(t *testing.T) {
	sac := New()
	k := "key"

	v := int64(math.MinInt64)
	sac.Set(k, v)

	n := sac.GetInt64(k)
	simpleAssert(t, n, v)

	nAny := sac.GetNumber(k)
	simpleAssert(t, nAny, int(v))
}

func TestRoundNumber32(t *testing.T) {
	sac := New()
	k := "key"

	v := float32(3.14)
	sac.Set(k, v)

	n := sac.GetNumber(k)
	simpleAssert(t, n, 3)
}

func TestRoundNumber64(t *testing.T) {
	sac := New()
	k := "key"

	v := float64(3.14)
	sac.Set(k, v)

	n := sac.GetNumber(k)
	simpleAssert(t, n, 3)
}
