package problems

func cmp(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func findAnagrams(s string, p string) []int {
	var res []int

	l := len(p)

	if l > len(s) {
		return res
	}

	pLetters := make([]int, 26)
	for _, ch := range p {
		pLetters[int(ch-'a')]++
	}

	winLetters := make([]int, 26)
	for i := 0; i < l; i++ {
		winLetters[int(s[i]-'a')]++
	}

	for i := 0; i <= len(s)-l; i++ {
		if i > 0 {
			winLetters[int(s[i-1]-'a')]--
			winLetters[int(s[i+l-1]-'a')]++
		}

		if cmp(winLetters, pLetters) {
			res = append(res, i)
		}
	}
	return res
}
