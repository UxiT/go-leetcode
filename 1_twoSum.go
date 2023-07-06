package algos

func twoSum(nums []int, target int) []int {
	var indexes []int

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < (len(nums) - i); j++ {
			if nums[i]+nums[j] == target {
				indexes = append(indexes, i, i)
			}
		}
	}

	return indexes
}
