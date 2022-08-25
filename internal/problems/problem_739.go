package problems

func dailyTemperatures(temperatures []int) []int {
	l := len(temperatures)
	res := make([]int, l)
	for i := l - 2; i >= 0; i-- {
		cur := temperatures[i]
		for j := i + 1; ; {
			if temperatures[j] > cur {
				res[i] = j - i
				break
			} else {
				if res[j] == 0 {
					res[i] = 0
					break
				} else {
					j += res[j]
				}
			}
		}
	}
	return res
}
