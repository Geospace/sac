package sac

// GetNumber reads a value from the store, if it is any number
// the return type being forced to int, the actual underlying value may
// overflow, or get truncated
func (s *Sac) GetNumber(k string) int {
	v := s.Get(k)

	if v == nil {
		return 0
	}

	switch v := v.(type) {
	case uint8:
		return int(v)
	case int8:
		return int(v)
	case uint16:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int:
		return v
		// -----------------
		// Fear the truncation
		// -----------------
	case float32:
		return int(v)
	case float64:
		return int(v)
		// -----------------
		// Fear the overflow
		// -----------------
	case uint:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v)
	case int64:
		return int(v)
	}

	return 0
}

// GetUint8 reads a value from the store, if it is an uint8
func (s *Sac) GetUint8(k string) uint8 {
	v := s.Get(k)

	if v == nil {
		return 0
	}

	if casted, ok := v.(uint8); ok {
		return casted
	}

	return 0
}

// GetUint16 reads a value from the store, if it is an uint16
func (s *Sac) GetUint16(k string) uint16 {
	v := s.Get(k)

	if v == nil {
		return 0
	}

	if casted, ok := v.(uint16); ok {
		return casted
	}

	return 0
}

// GetUint32 reads a value from the store, if it is an uint32
func (s *Sac) GetUint32(k string) uint32 {
	v := s.Get(k)

	if v == nil {
		return 0
	}

	if casted, ok := v.(uint32); ok {
		return casted
	}

	return 0
}

// GetUint64 reads a value from the store, if it is an uint64
func (s *Sac) GetUint64(k string) uint64 {
	v := s.Get(k)

	if v == nil {
		return 0
	}

	if casted, ok := v.(uint64); ok {
		return casted
	}

	return 0
}

// GetInt8 reads a value from the store, if it is an int8
func (s *Sac) GetInt8(k string) int8 {
	v := s.Get(k)

	if v == nil {
		return 0
	}

	if casted, ok := v.(int8); ok {
		return casted
	}

	return 0
}

// GetInt16 reads a value from the store, if it is an int16
func (s *Sac) GetInt16(k string) int16 {
	v := s.Get(k)

	if v == nil {
		return 0
	}

	if casted, ok := v.(int16); ok {
		return casted
	}

	return 0
}

// GetInt32 reads a value from the store, if it is an int32
func (s *Sac) GetInt32(k string) int32 {
	v := s.Get(k)

	if v == nil {
		return 0
	}

	if casted, ok := v.(int32); ok {
		return casted
	}

	return 0
}

// GetInt64 reads a value from the store, if it is an int64
func (s *Sac) GetInt64(k string) int64 {
	v := s.Get(k)

	if v == nil {
		return 0
	}

	if casted, ok := v.(int64); ok {
		return casted
	}

	return 0
}
