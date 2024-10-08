package medium

import (
	"strconv"
	"strings"
)

func CompareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	vInt1 := make([]int, len(v1))
	vInt2 := make([]int, len(v2))

	for i := 0; i < len(v1); i++ {
		vInt1[i], _ = strconv.Atoi(v1[i])
	}
	for i := 0; i < len(v2); i++ {
		vInt2[i], _ = strconv.Atoi(v2[i])
	}

	i := 0
	j := 0

	for i < len(vInt1) && j < len(vInt2) {

		if vInt1[i] < vInt2[j] {
			return -1
		} else if vInt1[i] > vInt2[j] {
			return 1
		}
		i++
		j++
	}

	for i < len(vInt1) {
		if vInt1[i] > 0 {
			return 1
		}
		i++
	}
	for j < len(vInt2) {
		if vInt2[j] > 0 {
			return -1
		}
		j++
	}
	return 0
}
