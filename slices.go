package sac

// GetStringSlice reads a value from the store, if it is a slice of strings
func (s *Sac) GetStringSlice(k string) []string {
	v := s.Get(k)

	if v == nil {
		return []string{}
	}

	if casted, ok := v.([]string); ok {
		return casted
	}

	return []string{}
}
