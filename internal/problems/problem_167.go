package problems

func twoSum(numbers []int, target int) []int {
	beg := 0
	end := len(numbers) - 1

	for beg != end {
		cur := numbers[beg] + numbers[end]
		if target > cur {
			beg++
		} else if target < cur {
			end--
		} else {
			return []int{beg + 1, end + 1}
		}
	}

	return []int{-1, -1}
}
