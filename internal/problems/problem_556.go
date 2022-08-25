package problems

import "sort"

func nextGreaterElement(n int) int {
	digits := make([]int, 0)
	for ; n > 0; n /= 10 {
		digits = append(digits, n%10)
	}

	st := make([]int, 0)
	size := len(digits)
	nextLess := make([]int, size)
	st = append(st, 0)
	for i := 1; i < size; i++ {
		for len(st) > 0 && digits[i] < digits[st[len(st)-1]] {
			nextLess[st[len(st)-1]] = i
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}

	for len(st) > 0 {
		nextLess[st[len(st)-1]] = -1
		st = st[:len(st)-1]
	}

	digitIdx := -1
	swapIdx := len(digits)
	for i := 0; i < size; i++ {
		if nextLess[i] == -1 {
			continue
		}
		if swapIdx > nextLess[i] {
			digitIdx = i
			swapIdx = nextLess[i]
		}
	}

	if digitIdx == -1 {
		return -1
	}

	digits[digitIdx], digits[swapIdx] = digits[swapIdx], digits[digitIdx]

	sort.Slice(digits[0:swapIdx], func(i, j int) bool { return digits[i] > digits[j] })

	mpl := 1
	sum := 0
	for i := 0; i < size; i++ {
		newSum := sum + mpl*digits[i]
		if int(int32(newSum)) < newSum {
			return -1
		}

		sum = newSum
		mpl *= 10
	}

	return sum
}
