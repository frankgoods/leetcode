package problems

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	sum := 0
	for i, v := range nums {
		sum += v
		nums[i] = sum
	}

	ma := nums[0]
	mi := 10001
	for _, v := range nums {
		ma = max(ma, max(v, v-mi))
		mi = min(v, mi)
	}

	return ma
}
