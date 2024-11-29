package hash

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := make(map[int32]uint)
	for _, charValue := range s {
		sMap[charValue] += 1
	}

	tMap := make(map[int32]uint)
	for _, charValue := range t {
		tMap[charValue] += 1
	}

	if len(sMap) != len(tMap) {
		return false
	}

	for k, v := range sMap {
		if tMap[k] != v {
			return false
		}
	}

	return true

}
