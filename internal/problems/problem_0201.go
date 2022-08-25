package problems

func rangeBitwiseAnd(left int, right int) int {
	res := 0
	cur := 1
	for {
		if left == 0 && right == 0 {
			return res
		}
		if (left == 0) || (right == 0) {
			return 0
		}

		if left&1 != right&1 {
			res &= cur
		} else if left&1 == 1 {
			res |= cur
		}

		left >>= 1
		right >>= 1
		cur <<= 1
	}
	return res
}
