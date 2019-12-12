package sac

// GetFloat reads a value from the store, if it is a floating point number
func (s *Sac) GetFloat(k string) float64 {
	v := s.Get(k)

	if v == nil {
		return 0
	}

	switch v := v.(type) {
	case float64:
		return v
	case float32:
		// Both types are compatible
		return float64(v)
	}

	return 0
}

// GetFloat reads a value from the store, if it is a float32
func (s *Sac) GetFloat32(k string) float32 {
	v := s.Get(k)

	if v == nil {
		return 0
	}

	if casted, ok := v.(float32); ok {
		return casted
	}

	return 0
}

// GetFloat reads a value from the store, if it is a float64
func (s *Sac) GetFloat64(k string) float64 {
	v := s.Get(k)

	if v == nil {
		return 0
	}

	if casted, ok := v.(float64); ok {
		return casted
	}

	return 0
}
