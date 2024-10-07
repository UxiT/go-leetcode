package algos

func StrStr(haystack string, needle string) int {
	haystackLen := len(haystack)
	needleLen := len(needle)

	if needleLen > haystackLen {
		return -1
	}

	if needle == haystack {
		return 0
	}

	for i := 0; i < (haystackLen - needleLen + 1); i++ {
		if haystack[i:i+needleLen] == needle {
			return i
		}
	}

	return -1
}
