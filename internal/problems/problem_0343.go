package problems

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func br(second bool, n int, memo []int) int {
	if second && n <= 4 {
		return n
	}

	if memo[n] != 0 {
		return memo[n]
	}

	m := 0
	for i := 1; i <= n/2+1; i++ {
		m = max(i*br(true, n-i, memo), m)
	}
	memo[n] = m
	return memo[n]
}

func integerBreak(n int) int {
	memo := make([]int, n+1)
	return br(false, n, memo)
}
