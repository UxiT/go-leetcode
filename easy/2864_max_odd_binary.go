package algos

import "strings"

func MaximumOddBinaryNumber(s string) string {
	zeroes := 0
	ones := 0

	for _, c := range s {
		if c == '1' {
			ones++
		} else {
			zeroes++
		}
	}

	return strings.Repeat("1", ones-1) + strings.Repeat("0", zeroes) + "1"
}
