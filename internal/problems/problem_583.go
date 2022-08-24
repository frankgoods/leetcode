package problems

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDist(w1, w2 string, i1, i2 int, memo [][]int) int {
	if i1 == len(w1) {
		return len(w2) - i2
	}
	if i2 == len(w2) {
		return len(w1) - i1
	}

	if memo[i1][i2] != -1 {
		return memo[i1][i2]
	}

	if w1[i1] == w2[i2] {
		memo[i1][i2] = minDist(w1, w2, i1+1, i2+1, memo)
	} else {
		memo[i1][i2] = 1 + min(minDist(w1, w2, i1+1, i2, memo), minDist(w1, w2, i1, i2+1, memo))
	}

	return memo[i1][i2]
}

func minDistance(word1 string, word2 string) int {
	l1 := len(word1)
	l2 := len(word2)
	memo := make([][]int, l1)
	for i := range memo {
		memo[i] = make([]int, l2)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return minDist(word1, word2, 0, 0, memo)
}
