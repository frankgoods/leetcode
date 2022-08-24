package problems

import "sort"

const empty = 100

func perm(nums []int, curSet []int, res *[][]int) {
	if len(curSet) == len(nums) {
		cpy := append([]int{}, curSet...)
		*res = append(*res, cpy)
		return
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == empty || (i > 0 && nums[i-1] == nums[i]) {
			continue
		}
		tmp := nums[i]
		nums[i] = empty
		perm(nums, append(curSet, tmp), res)
		nums[i] = tmp
	}
}

func permuteUnique(nums []int) [][]int {
	var res [][]int

	sort.Ints(nums)
	perm(nums, []int{}, &res)

	return res
}
