package problems

func numSubarrayProductLessThanK(nums []int, k int) int {
	result := 0

	mpl := 1
	curLen := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] >= k {
			mpl = 1
			curLen = 0
			continue
		}
		if curLen > 0 {
			mpl /= nums[i-1]
			curLen--
		}
		for i+curLen < len(nums) && mpl*nums[i+curLen] < k {
			mpl *= nums[i+curLen]
			curLen++
		}

		if i+curLen == len(nums) {
			result += int(float64(1+curLen) / 2.0 * float64(curLen))
			return result
		}

		result += curLen
	}
	return result
}
