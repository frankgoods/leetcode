package problems

func lengthOfLIS(nums []int) int {
	l := len(nums)
	dp := make([]int, l)
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
	}

	m := 0
	for i := range dp {
		if dp[i] > m {
			m = dp[i]
		}
	}

	return m
}
