package problems

func take(items []int, idx int, taken []int, result *[][]int, k int) {
	if len(taken) == k {
		tmp := make([]int, k)
		copy(tmp, taken)
		*result = append(*result, tmp)
		return
	}

	for i := idx; i < len(items); i++ {
		taken = append(taken, items[i])
		take(items, i+1, taken, result, k)
		taken = taken[:len(taken)-1]
	}
}

func combine(n int, k int) [][]int {
	taken := make([]int, 0, k)
	var result [][]int

	items := make([]int, n)
	for idx, _ := range items {
		items[idx] = idx + 1
	}

	take(items, 0, taken, &result, k)

	return result
}
