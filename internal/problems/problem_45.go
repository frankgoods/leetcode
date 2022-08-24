package problems

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func jump(nums []int) int {
	l := len(nums)
	jump := make([]int, l)
	for i := range jump {
		jump[i] = 1000000
	}
	jump[0] = 0

	for i := range nums {
		for j := 1; j <= nums[i]; j++ {
			if i+j < l {
				jump[i+j] = min(jump[i]+1, jump[i+j])
			}
		}
	}

	return jump[l-1]
}
