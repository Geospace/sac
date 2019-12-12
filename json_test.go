package sac

import (
	"testing"
)

const (
	extJSON = ".json"
)

func TestErrorJSON(t *testing.T) {
	testError(t, JSON, extJSON)
}

func TestReadJSON(t *testing.T) {
	testRead(t, JSON, extJSON)
}

func TestReadNestedJSON(t *testing.T) {
	testReadNested(t, JSON, extJSON)
}

func TestReadEmptyJSON(t *testing.T) {
	testReadEmpty(t, JSON, extJSON)
}

func TestReadNotExistJSON(t *testing.T) {
	testReadNotExist(t, JSON, extJSON)
}

func TestReadCaseSensitiveJSON(t *testing.T) {
	testReadCaseSensitive(t, JSON, extJSON)
}

func TestWriteMemory(t *testing.T) {
	testWriteMemory(t, JSON, extJSON)
}

func TestWriteFSSimpleJSON(t *testing.T) {
	testWriteFSSimple(t, JSON, extJSON)
}

func TestWriteFSSafeJSON(t *testing.T) {
	testWriteFSSafe(t, JSON, extJSON)
}
