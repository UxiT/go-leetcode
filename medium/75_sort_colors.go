package medium

func sortColors(nums []int) {
	firstElem := 0
	numsLen := len(nums) - 1

	for i := numsLen; i > firstElem; i-- {
		for nums[firstElem] == 0 && firstElem < numsLen {
			firstElem++
		}

		if nums[i] == 0 {
			nums[i] = nums[firstElem]
			nums[firstElem] = 0
		}
	}

	for i := numsLen; i > firstElem; i-- {
		for (nums[firstElem] == 1 || nums[firstElem] == 0) && firstElem < numsLen {
			firstElem++
		}

		if nums[i] == 1 {
			nums[i] = nums[firstElem]
			nums[firstElem] = 1
		}
	}
}
