// sac is a simple key/value store that can work with
// a configuration file on your filesystem
// no automatic conversions
// no error handling, returns the 0 equivalent
// case sensitive
package sac

import "strings"

type Sac struct {
	fields map[string]interface{}
}

const (
	nestingChar = "."
)

// New returns a new sac, ready to be used
func New() *Sac {
	fields := make(map[string]interface{})
	return &Sac{
		fields,
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
	level := getNest(s.fields, split[0:len(split)-1])

	level[k] = v
}

// GetString reads a value from the store, as a string
func (s *Sac) GetString(k string) string {
	split := strings.Split(k, nestingChar)
	level := getNest(s.fields, split[0:len(split)-1])

	if v, ok := level[k]; ok {
		if casted, okT := v.(string); okT {
			return casted
		}
	}

	return ""
}

// GetNumber reads a value from the store, as a number
// int is the only supported type
func (s *Sac) GetNumber(k string) int {
	split := strings.Split(k, nestingChar)
	level := getNest(s.fields, split[0:len(split)-1])

	if v, ok := level[k]; ok {
		if casted, okT := v.(int); okT {
			return casted
		}
	}

	return 0
}

// GetNumber reads a value from the store, as a boolean
func (s *Sac) GetBool(k string) bool {
	split := strings.Split(k, nestingChar)
	level := getNest(s.fields, split[0:len(split)-1])

	if v, ok := level[k]; ok {
		if casted, okT := v.(bool); okT {
			return casted
		}
	}

	return false
}
