package hash

import "testing"

func commonCharsNoDiff(words []string) []string {
	checkResult := len(words)
	wordMap := make(map[int32]int)
	for _, word := range words {
		firstWordCheckMap := make(map[int32]bool)
		for _, wordChar := range word {
			if _, ok := firstWordCheckMap[wordChar]; !ok {
				wordMap[wordChar] += 1
				firstWordCheckMap[wordChar] = true
			}
		}
	}

	result := make([]string, 0, len(wordMap))
	for wordChar, num := range wordMap {
		if num == checkResult {
			result = append(result, string(wordChar))
		}
	}

	return result
}

func commonChars(words []string) []string {
	wordChars := make([][]int32, 0, len(words))
	for _, word := range words {
		wordChar := make([]int32, 26)
		for _, ch := range word {
			wordChar[ch-'a'] += 1
		}

		wordChars = append(wordChars, wordChar)
	}

	n := len(words)

	result := make([]string, 0, n)
	for i := 0; i < 26; i++ {
		minNum := wordChars[0][i]
		for j := 1; j < n; j++ {
			if wordChars[j][i] < minNum {
				minNum = wordChars[j][i]
			}
		}

		for k := 0; k < int(minNum); k++ {
			result = append(result, string(rune('a'+i)))
		}
	}

	return result

}

func Test_1002(t *testing.T) {
	res := commonCharsNoDiff([]string{"bella", "label", "roller"})
	if len(res) != 2 {
		t.Error("commonChars failed")
	}

	res1 := commonCharsNoDiff([]string{"cool", "lock", "cook"})
	if len(res1) != 2 {
		t.Error("commonChars failed")
	}

	res2 := commonChars([]string{"bella", "label", "roller"})
	if len(res2) != 3 {
		t.Error("commonChars failed")
	}

	res3 := commonChars([]string{"cool", "lock", "cook"})
	if len(res3) != 2 {
		t.Error("commonChars failed")
	}
}
