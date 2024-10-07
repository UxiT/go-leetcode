package easy

// Assume that largest number is always on the left
func RomanToInt(s string) int {
	sum := 0

	symMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := range s {
		if (i+1) < len(s) && symMap[s[i+1]] > symMap[s[i]] {
			sum -= symMap[s[i]]
		} else {
			sum += symMap[s[i]]
		}
	}

	return sum
}

// Runtime 88.45%
// Memory 88.72%
