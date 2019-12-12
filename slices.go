package sac

// GetStringSlice reads a value from the store, if it is a slice of strings
func (s *Sac) GetStringSlice(k string) []string {
	v := s.Get(k)

	zeroValue := []string{}

	if v == nil {
		return zeroValue
	}

	switch v := v.(type) {
	case []string:
		return v
	case []interface{}:
		var strSlice []string

		for _, e := range v {
			if elementCasted, okk := e.(string); okk {
				strSlice = append(strSlice, elementCasted)
			} else {
				return zeroValue
			}
		}

		return strSlice
	}

	if casted, ok := v.([]string); ok {
		return casted
	}

	return zeroValue
}
