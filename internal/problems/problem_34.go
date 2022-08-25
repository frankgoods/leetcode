package problems

func bound(nums []int, target int, lower bool) int {
	if len(nums) < 1 {
		return -1
	}

	l := 0
	r := len(nums) - 1

	if lower && nums[l] == target {
		return l
	}

	if !lower && nums[r] == target {
		return r
	}

	for r-l > 1 {
		m := (r + l) / 2
		if lower {
			if nums[m] < target {
				l = m
			} else {
				r = m
			}
		} else {
			if nums[m] > target {
				r = m
			} else {
				l = m
			}
		}
	}

	if lower && nums[r] == target {
		return r
	} else if nums[l] == target {
		return l
	} else {
		return -1
	}
}

func searchRange(nums []int, target int) []int {
	return []int{bound(nums, target, true), bound(nums, target, false)}
}
