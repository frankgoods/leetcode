package problems

import "sort"

func threeSum(nums []int) [][]int {
	var res [][]int

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		for i > 0 && i < len(nums) && nums[i] == nums[i-1] {
			i++
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[l] + nums[r]
			if sum < -nums[i] {
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			} else if sum > -nums[i] {
				r--
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				r--
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}
		}
	}

	return res
}
