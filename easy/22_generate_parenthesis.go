package algos

import "strings"

func GenerateParenthesis(n int) []string {
	var result []string
	result = append(result, strings.Repeat("(", n)+strings.Repeat(")", n))

	if n <= 1 {
		return result
	}

	if n == 2 {
		result = append(result, "()()")
		return result
	}

	for i := n; i > 0; {
		var tmp string

		if i == n {
			tmp = strings.Repeat("(", i) + strings.Repeat(")", n)
			result = append(result, tmp)
			i -= 1
			continue
		}

	}

	return result
}

// PAUSED
