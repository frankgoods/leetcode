package problems

func rotate(nums []int, k int) {
	length := len(nums)
	k %= length

	count := 0
	for i := 0; count < length-1; i++ {
		hole := i
		val := nums[i]
		j := i
		for {
			j = (j + k) % length
			tmp := nums[j]
			nums[j] = val
			val = tmp
			count++
			if j == hole {
				break
			}
		}
	}
}
