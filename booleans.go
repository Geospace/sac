package sac

// GetBool reads a value from the store, if it is a boolean
func (s *Sac) GetBool(k string) bool {
	v := s.Get(k)

	if v == nil {
		return false
	}

	if casted, ok := v.(bool); ok {
		return casted
	}

	return false
}
