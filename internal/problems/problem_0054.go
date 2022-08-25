package problems

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])

	res := make([]int, 0, m*n)

	for i := 0; i < (m+1)/2 && i < (n+1)/2; i++ {
		walk(matrix, i, i, m-2*i, n-2*i, &res)
	}
	return res
}

func walk(mat [][]int, startI, startJ int, m, n int, res *[]int) {
	for j := startJ; j < startJ+n; j++ {
		*res = append(*res, mat[startI][j])
	}
	for i := startI + 1; i < startI+m; i++ {
		*res = append(*res, mat[i][startJ+n-1])
	}

	for j := startJ + n - 2; j >= startJ && m > 1; j-- {
		*res = append(*res, mat[startI+m-1][j])
	}
	for i := startI + m - 2; i > startI && n > 1; i-- {
		*res = append(*res, mat[i][startJ])
	}
}
