package problems

func nextGreaterElements(nums []int) []int {
	stack := []int{}
	n := len(nums)
	res := make([]int, n)
	cycle := false
	for i := 0; i < 2*n; i++ {
		if i == n {
			cycle = true
		}
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			idx := stack[len(stack)-1]
			res[idx] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		if !cycle {
			stack = append(stack, i)
		}
	}

	for len(stack) > 0 {
		idx := stack[len(stack)-1]
		res[idx] = -1
		stack = stack[:len(stack)-1]
	}
	return res
}
