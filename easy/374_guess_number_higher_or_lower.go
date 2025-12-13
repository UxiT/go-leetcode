package easy

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
var target = 2

func guess(n int) int {
	if n == target {
		return 0
	}

	if n > target {
		return -1
	}

	return 1
}

func GuessNumber(n int) int {
	minVal, maxVal, current := 1, n+1, n/2

	for {
		res := guess(current)

		if res == 0 {
			return current
		}

		if res == 1 {
			minVal = current
			current = int((minVal + maxVal) / 2)
		}

		if res == -1 {
			maxVal = current
			current = int((minVal + current) / 2)
		}
	}
}
