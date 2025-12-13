package easy

import "log"

func Search(nums []int, target int) int {
	start, current, end := 0, len(nums)/2, len(nums)

	for {
		if current == start && nums[current] != target {
			return -1
		}

		if nums[current] == target {
			return current
		} else if nums[current] > target {
			end = current
			current = int((start + end) / 2)
		} else if nums[current] < target {
			start = current
			current = int((start + end) / 2)
		}

		log.Printf("start: %d, end: %d, current: %d", start, end, current)
	}
}
