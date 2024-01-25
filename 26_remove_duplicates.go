package algos

func removeDuplicates(nums []int) int {
	seen := make(map[int]bool)
	cnt := 0
	r := []int{}

	for _, v := range nums {
		if seen[v] == true {
			continue
		}

		seen[v] = true
		cnt += 1
		r = append(r, v)
	}

	for k, _ := range nums {
		nums[k] = r[k]
	}
	return cnt
}
