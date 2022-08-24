package problems

func isHappy(n int) bool {
	m := make(map[int]bool)

	for {

		if n == 1 {
			return true
		}

		sum := 0
		for ; n > 0; n /= 10 {
			d := n % 10
			sum += d * d
		}

		if _, exists := m[sum]; exists {
			return false
		}

		m[sum] = true
		n = sum
	}
}
