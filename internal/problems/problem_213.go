package problems

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func robDrugdillers(a []int) int {
	l := len(a)
	if l > 2 {
		a[2] += a[0]
	}
	for i := 3; i < l; i++ {
		a[i] = max(a[i]+a[i-2], a[i]+a[i-3])
	}
	return max(a[l-1], a[l-2])
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	return max(
		robDrugdillers(append([]int{}, nums[1:]...)),
		robDrugdillers(append([]int{}, nums[:len(nums)-1]...)))
}
