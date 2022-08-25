package problems

func subarraySum(nums []int, k int) int {
	res := 0
	n := len(nums)

	if nums[0] == k {
		res++
	}

	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
		if nums[i] == k {
			res++
		}
	}

	left := make(map[int]int)
	left[nums[0]]++
	for i := 1; i < n; i++ {
		if v, ok := left[nums[i]-k]; ok {
			res += v
		}
		left[nums[i]]++
	}

	return res
}
