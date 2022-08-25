package problems

func findRotation(mat [][]int, target [][]int) bool {
	r := []bool{true, true, true, true}
	l := len(mat)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if mat[i][j] != target[i][j] {
				r[0] = false
			}
			if mat[i][j] != target[j][l-i-1] {
				r[1] = false
			}
			if mat[i][j] != target[l-i-1][l-1-j] {
				r[2] = false
			}
			if mat[i][j] != target[l-j-1][i] {
				r[3] = false
			}
		}
	}

	return r[0] || r[1] || r[2] || r[3]
}
