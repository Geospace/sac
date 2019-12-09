package sac

import (
	"path"
	"testing"
)

const (
	testMaterialFolder = "test_material"
)

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
	number := sac.GetNumber(k)
	simpleAssert(t, number, int64(v))

	vNoCast := int64(3)
	sac.Set(k, vNoCast)
	numberNoCast := sac.GetNumber(k)
	simpleAssert(t, numberNoCast, vNoCast)

	empty := sac.GetNumber("doesNotExist")
	simpleAssert(t, empty, int64(0))

	sac.Set(k, "this is wrong")
	wrong := sac.GetNumber(k)
	simpleAssert(t, wrong, int64(0))
}

func TestSetGetFloat(t *testing.T) {
	sac := New()

	k := "key"
	v := 3.14

	sac.Set(k, v)
	number := sac.GetFloat(k)
	simpleAssert(t, number, v)

	empty := sac.GetFloat("doesNotExist")
	simpleAssert(t, empty, float64(0))

	sac.Set(k, "this is wrong")
	wrong := sac.GetFloat(k)
	simpleAssert(t, wrong, float64(0))
}

func TestSetGetNumberDeep(t *testing.T) {
	sac := New()

	k := "some.nested.key"
	v := 3

	sac.Set(k, v)
	str := sac.GetNumber(k)
	simpleAssert(t, str, int64(v))

	empty := sac.GetNumber("does.not.exist")
	simpleAssert(t, empty, int64(0))

	sac.Set(k, "this is wrong")
	wrong := sac.GetNumber(k)
	simpleAssert(t, wrong, int64(0))
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

func TestErrorJSON(t *testing.T) {
	sac := New()

	pathNotExist := path.Join(testMaterialFolder, "does_not_exist.json")
	err := sac.ReadConfig(pathNotExist, JSON)
	trueAssert(t, err != nil)

	pathNoRead := path.Join(testMaterialFolder, "no_read.json")
	err = sac.ReadConfig(pathNoRead, JSON)
	trueAssert(t, err != nil)
}

func TestReadJSON(t *testing.T) {
	sac := New()

	path := path.Join(testMaterialFolder, "simple.json")
	err := sac.ReadConfig(path, JSON)
	errAssert(t, err)

	str := sac.GetString("key_string")
	simpleAssert(t, str, "string value")

	strOther := sac.GetString("key_string_other")
	simpleAssert(t, strOther, "other string value")

	number := sac.GetNumber("key_number")
	simpleAssert(t, number, int64(3))

	float := sac.GetFloat("key_float")
	simpleAssert(t, float, 3.14)

	boolean := sac.GetBool("key_bool")
	simpleAssert(t, boolean, true)
}

func TestReadJSONNested(t *testing.T) {
	sac := New()

	path := path.Join(testMaterialFolder, "nested.json")
	err := sac.ReadConfig(path, JSON)
	errAssert(t, err)

	str := sac.GetString("nested.key_string")
	simpleAssert(t, str, "string value")

	strOther := sac.GetString("very.nested.key_string_other")
	simpleAssert(t, strOther, "other string value")

	number := sac.GetNumber("very.very.nested.key_number")
	simpleAssert(t, number, int64(3))

	float := sac.GetFloat("key_float")
	simpleAssert(t, float, 3.14)

	boolean := sac.GetBool("nested.key_bool")
	simpleAssert(t, boolean, true)
}

func TestReadJSONEmpty(t *testing.T) {
	sac := New()

	path := path.Join(testMaterialFolder, "empty.json")
	err := sac.ReadConfig(path, JSON)
	errAssert(t, err)

	empty := sac.GetString("empty")
	simpleAssert(t, empty, "")
}

func TestReadJSONNotExist(t *testing.T) {
	sac := New()

	path := path.Join(testMaterialFolder, "empty.json")
	err := sac.ReadConfig(path, JSON)
	errAssert(t, err)

	empty := sac.GetString("DOES_NOT_EXIST")
	simpleAssert(t, empty, "")
}

func TestReadJSONCaseSensitive(t *testing.T) {
	sac := New()

	path := path.Join(testMaterialFolder, "case.json")
	err := sac.ReadConfig(path, JSON)
	errAssert(t, err)

	low := sac.GetString("key")
	simpleAssert(t, low, "value")

	up := sac.GetString("Key")
	simpleAssert(t, up, "VALUE")
}

func TestWriteJSONMemory(t *testing.T) {
	sac := New()

	path := path.Join(testMaterialFolder, "nested.json")
	err := sac.ReadConfig(path, JSON)
	errAssert(t, err)

	k := "very.nested.key_string_other"

	before := sac.GetString(k)
	simpleAssert(t, before, "other string value")

	sac.Set(k, "changed")

	after := sac.GetString(k)
	simpleAssert(t, after, "changed")
}
