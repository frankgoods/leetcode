package problems

func uniquePaths(m int, n int) int {
	paths := make([]int, n)
	for i := range paths {
		paths[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			paths[j] += paths[j-1]
		}
	}

	return paths[n-1]
}
