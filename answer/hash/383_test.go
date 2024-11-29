package hash

func canConstruct(ransomNote string, magazine string) bool {
	magazimeMap := make(map[int32]int)

	for _, char := range magazine {
		magazimeMap[char] += 1
	}

	for _, char := range ransomNote {
		if value, ok := magazimeMap[char]; ok && value > 0 {
			magazimeMap[char] -= 1

			continue
		}
		return false
	}

	return true
}
