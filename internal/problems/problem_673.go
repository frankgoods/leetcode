package problems

func findNumberOfLIS(nums []int) int {
	l := len(nums)
	dp := make([]int, l)
	cnt := make([]int, l)
	for i := range dp {
		dp[i] = 1
		cnt[i] = 1
	}

	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					cnt[i] = cnt[j]
				} else if dp[j]+1 == dp[i] {
					cnt[i] += cnt[j]
				}
			}
		}
	}

	m := 0
	c := 0
	for i := range dp {
		if dp[i] > m {
			m = dp[i]
			c = cnt[i]
		} else if dp[i] == m {
			c += cnt[i]
		}
	}

	return c
}
