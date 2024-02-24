package algos

func calcS(idx1, idx2, h1, h2 int) int {
	if h1 < h2 {
		return h1 * (idx2 - idx1)
	} else {
		return h2 * (idx2 - idx1)
	}
}

func maxArea(height []int) int {
	h1, h2 := 0, len(height)-1
	S := 0

	for h1 < h2 {
		newS := calcS(h1, h2, height[h1], height[h2])

		if newS > S {
			S = newS
		}

		if height[h1] < height[h2] {
			h1++
		} else {
			h2--
		}
	}

	return S
}
