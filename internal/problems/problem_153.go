package problems

func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1

	if nums[l] <= nums[r] {
		return nums[l]
	}

	for r-l > 1 {
		m := (r + l) / 2
		if nums[m] > nums[l] {
			l = m
		} else if nums[m] < nums[r] {
			r = m
		}
	}
	return nums[r]
}
