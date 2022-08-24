package problems

import "sort"

func backTrack(nums []int, i int, curSet []int, res *[][]int) {
	if i == len(nums) {
		cpy := append([]int{}, curSet...)
		*res = append(*res, cpy)
		return
	}

	dupNum := 0
	ii := i
	for ii < len(nums)-1 && nums[ii] == nums[ii+1] {
		dupNum++
		ii++
	}

	backTrack(nums, i+dupNum+1, curSet, res)

	for x := i; x < i+dupNum+1; x++ {
		curSet = append(curSet, nums[x])
		backTrack(nums, i+dupNum+1, curSet, res)
	}
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	var res [][]int
	backTrack(nums, 0, []int{}, &res)

	return res
}
