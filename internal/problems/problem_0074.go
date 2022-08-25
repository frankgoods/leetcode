package problems

func checkInside(row []int, target int) bool {
	size := len(row)
	if target == row[size-1] {
		return true
	}

	l := 0
	r := size - 1
	mid := (r + l) / 2
	for l != mid {
		if target > row[mid] {
			l = mid
		} else if target < row[mid] {
			r = mid
		} else {
			return true
		}
		mid = (l + r) / 2
	}

	return row[l] == target
}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	if target >= matrix[m-1][0] {
		return checkInside(matrix[m-1], target)
	}

	l := 0
	r := m - 1
	mid := (r + l) / 2
	for l != mid {
		if target < matrix[mid][0] {
			r = mid
		} else if target > matrix[mid][n-1] {
			l = mid
		} else {
			return checkInside(matrix[mid], target)
		}
		mid = (r + l) / 2
	}

	return checkInside(matrix[l], target)
}
