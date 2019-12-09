// sac is a simple key/value store that can work with
// a configuration file on your filesystem
// supports nested values (like.this.key for example)
// no automatic conversions
// no error handling by default, returns the 0 equivalent
// safe functions equivalents in case you need error handling
// case sensitive
package sac

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

// GetString reads a value from the store, as a string
func (s *Sac) GetString(k string) string {
	split := strings.Split(k, nestingChar)
	last := split[len(split)-1]
	level := getNest(s.fields, split[0:len(split)-1])

	if v, ok := level[last]; ok {
		if casted, okT := v.(string); okT {
			return casted
		}
	}

	return ""
}

// GetNumber reads a value from the store, as a number
// int is the only supported type
func (s *Sac) GetNumber(k string) int64 {
	split := strings.Split(k, nestingChar)
	last := split[len(split)-1]
	level := getNest(s.fields, split[0:len(split)-1])

	if v, ok := level[last]; ok {
		switch toCast := v.(type) {
		case int64:
			return toCast
		case float64:
			return int64(toCast)
		case int:
			return int64(toCast)
		}
	}

	return 0
}

// GetFloat reads a value from the store, as a floating number
// float64 is the only supported type
func (s *Sac) GetFloat(k string) float64 {
	split := strings.Split(k, nestingChar)
	last := split[len(split)-1]
	level := getNest(s.fields, split[0:len(split)-1])

	if v, ok := level[last]; ok {
		if casted, okT := v.(float64); okT {
			return casted
		}
	}

	return 0
}

// GetNumber reads a value from the store, as a boolean
func (s *Sac) GetBool(k string) bool {
	split := strings.Split(k, nestingChar)
	last := split[len(split)-1]
	level := getNest(s.fields, split[0:len(split)-1])

	if v, ok := level[last]; ok {
		if casted, okT := v.(bool); okT {
			return casted
		}
	}

	return false
}

func (s *Sac) readConfigJSON(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	jsonContent, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonContent, &s.fields)
}

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
