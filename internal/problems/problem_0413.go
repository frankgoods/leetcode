package problems

func numberOfArithmeticSlices(nums []int) int {
	sum := 0
	l := len(nums)
	for i := 0; i < l-2; {
		j := 0
		for d := nums[i+1] - nums[i]; i+j+2 < l && nums[i+j+2]-nums[i+j+1] == d; {
			j++
			sum += j
		}
		if j > 0 {
			i = i + j + 1
		} else {
			i++
		}
	}
	return sum
}
