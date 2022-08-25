package problems

import "math"

func getSubset(nums []int, val int) []int {
	res := make([]int, 0)

	i := 0
	for val != 0 {
		if val&1 == 1 {
			res = append(res, nums[i])
		}
		i++
		val >>= 1
	}

	return res
}

func subsets(nums []int) [][]int {
	res := make([][]int, 0)

	n := int(math.Pow(2.0, float64(len(nums))))
	for i := 0; i < n; i++ {
		res = append(res, getSubset(nums, i))
	}
	return res
}
