package sac

// GetString reads a value from the store, if it is a string
func (s *Sac) GetString(k string) string {
	v := s.Get(k)

	if v == nil {
		return ""
	}

	if casted, ok := v.(string); ok {
		return casted
	}

	return ""
}
