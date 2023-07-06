package algos

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		var min int

		if len(prefix) > len(strs[i]) {
			min = len(strs[i])
		} else {
			min = len(prefix)
		}

		for j := 0; j < min; j++ {
			if prefix[j] != strs[i][j] {
				prefix = prefix[:j]
				break
			}

		}
	}

	return prefix
}

// runtime 58.18%
// memory 33.96%
