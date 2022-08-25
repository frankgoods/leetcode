package problems

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	m := len(l)
	res := make([]bool, 0, m)

	for i := 0; i < m; i++ {
		res = append(res, isArith(nums, l[i], r[i]))
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isArith(nums []int, l, r int) bool {
	minV := nums[l]
	maxV := nums[l]

	for _, v := range nums[l : r+1] {
		minV = min(v, minV)
		maxV = max(v, maxV)
	}

	size := r - l + 1
	if (maxV-minV)%(size-1) != 0 {
		return false
	}

	d := (maxV - minV) / (size - 1)
	if d == 0 {
		return true
	}

	pos := make([]bool, size)
	for i := l; i <= r; i++ {
		if (nums[i]-minV)%d != 0 {
			return false
		}
		p := (nums[i] - minV) / d
		if pos[p] {
			return false
		}
		pos[p] = true
	}

	for _, v := range pos {
		if !v {
			return false
		}
	}
	return true
}
