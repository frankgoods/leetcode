package problems

func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)
	nums[0] %= k
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
		nums[i] %= k
		if nums[i] == 0 {
			return true
		}
	}

	left := make(map[int]bool)
	left[nums[0]] = true
	for i := 2; i < n; i++ {
		if _, ok := left[nums[i]]; ok {
			return true
		}
		left[nums[i-1]] = true
	}
	return false
}
