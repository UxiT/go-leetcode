package algos

func removeElement(nums []int, val int) int {
	l := len(nums)
	nextSwap := l - 1

	for i := nextSwap; i > 0; i-- {
		if i != val {
			nextSwap = i
			break
		}
	}

	for i := 0; i < l; i++ {
		if nums[i] == val {
			nums[i] = nums[nextSwap]
			nextSwap--
		}
	}

	nums = nums[:nextSwap]
	return l - nextSwap + 1
}
