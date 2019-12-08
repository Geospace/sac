// sac is a simple key/value store that can work with
// a configuration file on your filesystem
// no automatic conversions
// no error handling, returns the 0 equivalent
package sac

type Sac struct {
	fields map[string]interface{}
}

// New returns a new sac, ready to be used
func New() *Sac {
	fields := make(map[string]interface{})
	return &Sac{
		fields,
	}
}

// Set writes a value into the store
func (s *Sac) Set(k string, v interface{}) {
	s.fields[k] = v
}

// GetString reads a value from the store, as a string
func (s *Sac) GetString(k string) string {
	if v, ok := s.fields[k]; ok {
		if casted, okT := v.(string); okT {
			return casted
		}
	}

	return ""
}

// GetNumber reads a value from the store, as a number
// int is the only supported type
func (s *Sac) GetNumber(k string) int {
	if v, ok := s.fields[k]; ok {
		if casted, okT := v.(int); okT {
			return casted
		}
	}

	return 0
}

// GetNumber reads a value from the store, as a boolean
func (s *Sac) GetBool(k string) bool {
	if v, ok := s.fields[k]; ok {
		if casted, okT := v.(bool); okT {
			return casted
		}
	}

	return false
}
