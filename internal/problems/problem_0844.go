package problems

func backspaceCompare(s string, t string) bool {
	p1 := len(s) - 1
	p2 := len(t) - 1

	skipP1 := 0
	skipP2 := 0
	for p1 >= 0 || p2 >= 0 {
		if p1 >= 0 {
			if s[p1] == '#' {
				p1--
				skipP1++
				continue
			}
			if skipP1 > 0 {
				p1--
				skipP1--
				continue
			}
		}
		if p2 >= 0 {
			if t[p2] == '#' {
				p2--
				skipP2++
				continue
			}
			if skipP2 > 0 {
				p2--
				skipP2--
				continue
			}
		}

		if p1 == -1 || p2 == -1 {
			return false
		}

		if s[p1] != t[p2] {
			return false
		}
		p1--
		p2--
	}

	return true
}
