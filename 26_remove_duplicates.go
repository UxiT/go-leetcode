package algos

func removeDuplicates(nums []int) int {
	m := make(map[int]bool)
	k, c := 0, 0
	l := len(nums)

	for i := 0; i < l; i++ {
		if !m[nums[i]] {
			m[nums[i]] = true
			c++
			continue
		} else {
			k++
			nums = append(nums[:i], nums[i+1:]...)
		}
	}

	for i := 0; i < k; i++ {
		nums[l-i] = '_'
	}

	return c
}
