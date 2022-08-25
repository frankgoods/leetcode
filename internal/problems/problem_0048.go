package problems

func rotate(matrix [][]int) {
	m := len(matrix)

	for i := 0; i <= m/2; i++ {
		sqrLen := m - i*2 - 1
		for j := 0; j < sqrLen; j++ {
			x1 := i
			y1 := i // +

			x2 := i // +
			y2 := i + sqrLen
			tmp := matrix[x2+j][y2]
			matrix[x2+j][y2] = matrix[x1][y1+j]

			x3 := i + sqrLen
			y3 := i + sqrLen // -
			tmp, matrix[x3][y3-j] = matrix[x3][y3-j], tmp

			x4 := i + sqrLen // -
			y4 := i
			tmp, matrix[x4-j][y4] = matrix[x4-j][y4], tmp

			matrix[x1][y1+j] = tmp
		}
	}
}
