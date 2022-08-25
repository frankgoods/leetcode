package problems

func check(s string, pos int, size int) bool {
	for i := 0; i < size; i++ {
		if s[i] != s[i+pos] {
			return false
		}
	}

	return true
}

func repeatedSubstringPattern(s string) bool {
	l := len(s)
	if l < 2 {
		return false
	}

	for i := 1; i <= l/2; i++ {
		if l%i != 0 {
			continue
		}
		j := i
		end := l - i
		for ; j <= end; j += i {
			if !check(s, j, i) {
				break
			}
		}
		if j > end {
			return true
		}
	}

	return false
}
