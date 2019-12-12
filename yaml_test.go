package sac

import (
	"testing"
)

const (
	extYAML = ".yml"
)

func TestErrorYAML(t *testing.T) {
	testError(t, YAML, extYAML)
}

func TestReadYAML(t *testing.T) {
	testRead(t, YAML, extYAML)
}

func TestReadNestedYAML(t *testing.T) {
	testReadNested(t, YAML, extYAML)
}

func TestReadEmptyYAML(t *testing.T) {
	testReadEmpty(t, YAML, extYAML)
}

func TestReadNotExistYAML(t *testing.T) {
	testReadNotExist(t, YAML, extYAML)
}

func TestReadCaseSensitiveYAML(t *testing.T) {
	testReadCaseSensitive(t, YAML, extYAML)
}

func TestWriteMemoryYAML(t *testing.T) {
	testWriteMemory(t, YAML, extYAML)
}

func TestWriteFSSimpleYAML(t *testing.T) {
	testWriteFSSimple(t, YAML, extYAML)
}

func TestWriteFSSafeYAML(t *testing.T) {
	testWriteFSSafe(t, YAML, extYAML)
}
