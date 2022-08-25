package problems

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dist(w1, w2 string, i1, i2 int, memo [][]int) int {
	if i2 == len(w2) {
		return len(w1) - i1
	}
	if i1 == len(w1) {
		return len(w2) - i2
	}

	if memo[i1][i2] != -1 {
		return memo[i1][i2]
	}

	if w1[i1] == w2[i2] {
		return dist(w1, w2, i1+1, i2+1, memo)
	}

	f1 := 1 + dist(w1, w2, i1, i2+1, memo)
	f2 := 1 + dist(w1, w2, i1+1, i2, memo)
	f3 := 1 + dist(w1, w2, i1+1, i2+1, memo)

	memo[i1][i2] = min(min(f1, f2), f3)
	return memo[i1][i2]
}

func minDistance(word1 string, word2 string) int {
	memo := make([][]int, len(word1))
	for i := range memo {
		memo[i] = make([]int, len(word2))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return dist(word1, word2, 0, 0, memo)
}
