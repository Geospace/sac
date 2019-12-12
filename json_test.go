package sac

import (
	"io/ioutil"
	"path"
	"testing"

	"github.com/spf13/afero"
)

func TestJSONSetup(t *testing.T) {
	sac := New()

	fpath := path.Join(testMaterialFolder, "simple.json")
	err := sac.ReadConfig(fpath, JSON)
	errAssert(t, err)

	simpleAssert(t, sac.ConfigType, JSON)
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

	fpath := path.Join(testMaterialFolder, "simple.json")
	err := sac.ReadConfig(fpath, JSON)
	errAssert(t, err)

	str := sac.GetString("key_string")
	simpleAssert(t, str, "string value")

	strOther := sac.GetString("key_string_other")
	simpleAssert(t, strOther, "other string value")

	number := sac.GetNumber("key_number")
	simpleAssert(t, number, 3)

	float := sac.GetFloat("key_float")
	simpleAssert(t, float, 3.14)

	boolean := sac.GetBool("key_bool")
	simpleAssert(t, boolean, true)
}

func TestReadJSONNested(t *testing.T) {
	sac := New()

	fpath := path.Join(testMaterialFolder, "nested.json")
	err := sac.ReadConfig(fpath, JSON)
	errAssert(t, err)

	str := sac.GetString("nested.key_string")
	simpleAssert(t, str, "string value")

	strOther := sac.GetString("very.nested.key_string_other")
	simpleAssert(t, strOther, "other string value")

	number := sac.GetNumber("very.very.nested.key_number")
	simpleAssert(t, number, 3)

	float := sac.GetFloat("key_float")
	simpleAssert(t, float, 3.14)

	boolean := sac.GetBool("nested.key_bool")
	simpleAssert(t, boolean, true)
}

func TestReadJSONEmpty(t *testing.T) {
	sac := New()

	fpath := path.Join(testMaterialFolder, "empty.json")
	err := sac.ReadConfig(fpath, JSON)
	errAssert(t, err)

	empty := sac.GetString("empty")
	simpleAssert(t, empty, "")
}

func TestReadJSONNotExist(t *testing.T) {
	sac := New()

	fpath := path.Join(testMaterialFolder, "empty.json")
	err := sac.ReadConfig(fpath, JSON)
	errAssert(t, err)

	empty := sac.GetString("DOES_NOT_EXIST")
	simpleAssert(t, empty, "")
}

func TestReadJSONCaseSensitive(t *testing.T) {
	sac := New()

	fpath := path.Join(testMaterialFolder, "case.json")
	err := sac.ReadConfig(fpath, JSON)
	errAssert(t, err)

	low := sac.GetString("key")
	simpleAssert(t, low, "value")

	up := sac.GetString("Key")
	simpleAssert(t, up, "VALUE")
}

func TestWriteJSONMemory(t *testing.T) {
	sac := New()

	fpath := path.Join(testMaterialFolder, "nested.json")
	err := sac.ReadConfig(fpath, JSON)
	errAssert(t, err)

	k := "very.nested.key_string_other"

	before := sac.GetString(k)
	simpleAssert(t, before, "other string value")

	sac.Set(k, "changed")

	after := sac.GetString(k)
	simpleAssert(t, after, "changed")
}

func TestWriteJSONFSSimple(t *testing.T) {
	sac := New()

	fpath := path.Join(testMaterialFolder, "simple.json")
	err := sac.ReadConfig(fpath, JSON)
	errAssert(t, err)

	sac.Set("key_string", "string value changed")
	sac.Set("key_string_new", "new string value")
	sac.Set("key_number", 3.14)
	sac.Delete("key_string_other")
	sac.Delete("key_float")

	fs := afero.NewMemMapFs()
	sac.ChangeFS(fs)
	err = sac.WriteConfig()
	errAssert(t, err)

	sacOutFile, err := fs.Open(fpath)
	errAssert(t, err)
	sacOut, err := ioutil.ReadAll(sacOutFile)
	errAssert(t, err)

	goodOut, err := ioutil.ReadFile(path.Join(testMaterialFolder, "write_simple.json"))
	errAssert(t, err)

	simpleAssert(t, string(sacOut), string(goodOut))
}
