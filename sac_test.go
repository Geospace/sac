package sac

import (
	"testing"
)

const (
	testMaterialFolder = "test_material"
)

func TestNew(t *testing.T) {
	sac := New()
	trueAssert(t, sac != nil)
}

func TestSetGetRaw(t *testing.T) {
	sac := New()

	k := "key"
	v := make([]uint8, 1)

	sac.Set(k, v)
	ptr := sac.Get(k)
	trueAssert(t, ptr != nil)
}

func TestGetNotExist(t *testing.T) {}
