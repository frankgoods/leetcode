package problems

import "sort"

func comb(cans []int, target, sum, start int, curSet []int, res *[][]int) {
	if sum == target {
		cpy := append([]int{}, curSet...)
		*res = append(*res, cpy)
		return
	} else if sum > target {
		return
	}

	for i := start; i < len(cans); i++ {
		if i > start && cans[i] == cans[i-1] {
			continue
		}
		comb(cans, target, sum+cans[i], i+1, append(curSet, cans[i]), res)
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	sort.Ints(candidates)
	comb(candidates, target, 0, 0, []int{}, &res)
	return res
}
