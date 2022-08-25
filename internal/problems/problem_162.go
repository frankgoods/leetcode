package problems

func findPeakElement(nums []int) int {
	l := 0
	r := len(nums) - 1

	for r-l > 1 {
		m := (r + l) / 2
		if nums[m] > nums[m-1] {
			l = m
		} else {
			r = m
		}
	}

	if nums[r] > nums[l] {
		return r
	} else {
		return l
	}
}
