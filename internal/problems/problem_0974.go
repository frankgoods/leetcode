package problems

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func subarraysDivByK(nums []int, k int) int {
	res := 0
	n := len(nums)

	nums[0] %= k
	if nums[0] == 0 {
		res++
	}
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
		nums[i] %= k
		if nums[i] == 0 {
			res++
		}
	}

	left := make(map[int]int)
	left[nums[0]]++
	for i := 1; i < n; i++ {
		if v, ok := left[nums[i]]; ok {
			res += v
		}
		if v, ok := left[nums[i]-k]; ok {
			res += v
		}
		if v, ok := left[nums[i]+k]; ok {
			res += v
		}
		left[nums[i]]++
	}
	return res
}
