func searchMatrix(matrix [][]int, target int) bool {
	l, r := 0, len(matrix)-1
	row := 0
	for l <= r {
		row = (l + r) / 2
		if matrix[row][0] == target {
			return true
		}
		if matrix[row][0] > target {
			r = row - 1
			continue
		} else if matrix[row][len(matrix[row])-1] < target {
			l = row + 1
			continue
		}
		break
	}

	l, r = 0, len(matrix[row])-1
	for l <= r {
		mid := (l + r) / 2
		if matrix[row][mid] == target {
			return true
		}
		if matrix[row][mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}