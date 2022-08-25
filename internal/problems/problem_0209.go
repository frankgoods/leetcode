package problems

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minSubArrayLen(target int, nums []int) int {
	l := 0
	r := 0
	sum := 0
	mi := len(nums) + 1
	for {
		if sum < target && r < len(nums) {
			r++
			sum += nums[r-1]
			continue
		}
		if sum >= target {
			ln := r - l
			if ln == 1 {
				return 1
			}
			mi = min(mi, ln)
			sum -= nums[l]
			l++
			continue
		}
		if r == len(nums) {
			break
		}
	}

	if mi > len(nums) {
		return 0
	} else {
		return mi
	}
}
