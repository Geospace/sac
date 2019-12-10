// sac is a simple key/value store that can work with
// a configuration file on your filesystem
// supports nested values (like.this.key for example)
// no automatic conversions
// no error handling by default, returns the 0 equivalent
// safe functions equivalents in case you need error handling
// case sensitive
package sac

import (
	"fmt"
	"strings"
)

type Sac struct {
	fields     map[string]interface{}
	configType ConfigType
}

const (
	nestingChar = "."
)

// New returns a new sac, ready to be used
func New() *Sac {
	fields := make(map[string]interface{})
	return &Sac{
		fields,
		UNKNOWN,
	}
}

func getNest(root map[string]interface{}, levels []string) map[string]interface{} {
	// No more levels, stop recursion here
	if len(levels) == 0 {
		return root
	}

	// Advance one level
	nextLevel := levels[1:]
	// Check if the level exist
	v, ok := root[levels[0]]

	if !ok {
		// This level is accessed but does not exist, so create the level before
		// we continue
		root[levels[0]] = make(map[string]interface{})
		casted, _ := root[levels[0]].(map[string]interface{})
		return getNest(casted, nextLevel)
	}

	// The level already existed, we can continue
	casted, _ := v.(map[string]interface{})
	return getNest(casted, nextLevel)
}

// Set writes a value into the store
func (s *Sac) Set(k string, v interface{}) {
	split := strings.Split(k, nestingChar)
	last := split[len(split)-1]
	level := getNest(s.fields, split[0:len(split)-1])

	level[last] = v
}

// Get reads a value from the store, without making any
// assertion about its type
func (s *Sac) Get(k string) interface{} {
	split := strings.Split(k, nestingChar)
	last := split[len(split)-1]
	level := getNest(s.fields, split[0:len(split)-1])

	if v, ok := level[last]; ok {
		return v
	}

	return nil
}

// ReadConfig reads the configuration file at path into the sac object,
// according to the format parameter
func (s *Sac) ReadConfig(path string, configType ConfigType) error {
	switch configType {
	case YAML:
		return fmt.Errorf("not implemented")
	case JSON:
		return s.readConfigJSON(path)
	default:
		return fmt.Errorf("bad config type")
	}
}
