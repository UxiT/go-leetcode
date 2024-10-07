package algos

func sortColors(nums []int) {
	countMap := make(map[int]int, 3)
	offset := 0

	for _, v := range nums {
		_, ok := countMap[v]

		if ok {
			countMap[v] += 1
		} else {
			countMap[v] = 1
		}
	}

	for i := 0; i < countMap[0]; i++ {
		nums[offset] = 0
		offset++
	}

	for i := 0; i < countMap[1]; i++ {
		nums[offset] = 1
		offset++
	}

	for i := 0; i < countMap[2]; i++ {
		nums[offset] = 2
		offset++
	}
}
