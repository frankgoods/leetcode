package problems

func findStart(nums []int) int {
	l := 0
	r := len(nums) - 1

	if nums[0] <= nums[r] {
		return 0
	}

	for r-l > 1 {
		m := (r + l) / 2
		if nums[m] > nums[l] {
			l = m
		} else if nums[m] < nums[r] {
			r = m
		}
	}

	return l + 1
}

func findVal(nums []int, target int, idx func(int) int) int {
	l := 0
	r := len(nums) - 1

	if nums[idx(r)] == target {
		return idx(r)
	}
	if nums[idx(l)] == target {
		return idx(l)
	}

	for r-l > 1 {
		m := (r + l) / 2
		if nums[idx(m)] > target {
			r = m
		} else if nums[idx(m)] < target {
			l = m
		} else {
			return idx(m)
		}
	}

	return -1
}

func search(nums []int, target int) int {
	st := findStart(nums)
	return findVal(nums, target, func(idx int) int { return (idx + st) % len(nums) })
}
