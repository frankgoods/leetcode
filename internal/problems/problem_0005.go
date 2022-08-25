package problems

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func check(a, b int, s string) (int, int) {
	if b >= len(s) || s[a] != s[b] {
		return a, b
	}

	for a >= 0 && b < len(s) && s[a] == s[b] {
		a--
		b++
	}
	return a + 1, b
}

func longestPalindrome(s string) string {
	l := len(s)
	m := 0
	ans := ""
	for i := 0; i < l; i++ {

		a1, b1 := check(i, i, s)
		a2, b2 := check(i, i+1, s)
		if b1-a1 > m {
			m = b1 - a1
			ans = s[a1:b1]
		}
		if b2-a2 > m {
			m = b2 - a2
			ans = s[a2:b2]
		}
	}

	return ans
}
