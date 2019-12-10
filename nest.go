package sac

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
