package problems

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lcs(s1, s2 string, i1, i2 int, memo [][]int) int {
	if i1 == len(s1) || i2 == len(s2) {
		return 0
	}
	if memo[i1][i2] != -1 {
		return memo[i1][i2]
	}

	if s1[i1] == s2[i2] {
		memo[i1][i2] = 1 + lcs(s1, s2, i1+1, i2+1, memo)
	} else {
		memo[i1][i2] = max(lcs(s1, s2, i1+1, i2, memo), lcs(s1, s2, i1, i2+1, memo))
	}
	return memo[i1][i2]
}

func longestCommonSubsequence(text1 string, text2 string) int {
	l1 := len(text1)
	l2 := len(text2)
	memo := make([][]int, l1)
	for i := range memo {
		memo[i] = make([]int, l2)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return lcs(text1, text2, 0, 0, memo)
}
