package medium

// SearchMatrix results: 100% runtime, 86,18% memory
func SearchMatrix(matrix [][]int, target int) bool {
	start, end, row, column := 0, len(matrix), len(matrix)/2, len(matrix[0])/2

	for {
		if row == start && !(matrix[row][0] <= target && matrix[row][len(matrix[row])-1] >= target) {
			return false
		}

		if matrix[row][0] <= target && matrix[row][len(matrix[row])-1] >= target {
			break
		}

		if matrix[row][0] > target {
			end = row
			row = (start + end) / 2
		}

		if matrix[row][0] < target {
			start = row
			row = (start + end) / 2
		}
	}

	start = 0
	end = len(matrix[row])

	for {
		if column == start && matrix[row][column] != target {
			return false
		}

		if matrix[row][column] == target {
			return true
		}

		if matrix[row][column] > target {
			end = column
			column = (start + end) / 2
		}

		if matrix[row][column] < target {
			start = column
			column = (start + end) / 2
		}
	}
}
