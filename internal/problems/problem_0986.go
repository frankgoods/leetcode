package problems

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

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var res [][]int

	for i, j := 0, 0; i < len(firstList); i++ {
		if j > 0 {
			j--
		}

		if j == len(secondList)-1 && secondList[j][1] < firstList[i][0] {
			break
		}

		for j < len(secondList) && secondList[j][0] <= firstList[i][1] {
			if secondList[j][1] >= firstList[i][0] {
				res = append(res, []int{max(firstList[i][0], secondList[j][0]), min(firstList[i][1], secondList[j][1])})
			}
			j++
		}
	}

	return res
}
