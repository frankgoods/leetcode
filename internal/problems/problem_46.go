package problems

func take(nums []int, taken []int, result *[][]int) {
	if len(taken) == len(nums) {
		tmp := make([]int, len(taken))
		copy(tmp, taken)
		*result = append(*result, tmp)
	}

	const Empty = -11
	for idx, v := range nums {
		if v != Empty {
			tmp := v
			nums[idx] = Empty
			taken = append(taken, tmp)
			take(nums, taken, result)

			nums[idx] = tmp
			taken = taken[0 : len(taken)-1]
		}
	}
}

func permute(nums []int) [][]int {
	taken := make([]int, 0, len(nums))
	var result [][]int

	take(nums, taken, &result)

	return result
}
