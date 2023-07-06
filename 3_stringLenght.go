package algos

func lengthOfLongestSubstring(s string) int {
	size, maxSize := 0, 0
	chars := make(map[rune]bool)

	for j := 0; len(s)-j < maxSize; j++ {
		for i, v := range s[j:] {
			if chars[v] {
				if size > maxSize {
					maxSize = size
				}
				size = 0
				chars = make(map[rune]bool)
			}
			chars[v] = true
			size += 1

			if len(s) == i+1 && size > maxSize {
				maxSize = size
			}
		}
	}

	return maxSize
}
