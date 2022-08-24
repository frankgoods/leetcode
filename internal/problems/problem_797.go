package problems

func allPathsSourceTarget(graph [][]int) [][]int {
	var res [][]int
	dps(graph, 0, len(graph)-1, &res, []int{})

	return res
}

func dps(graph [][]int, i, n int, res *[][]int, curPath []int) {
	curPath = append(curPath, i)

	if i == n {
		cpy := append([]int{}, curPath...)
		*res = append(*res, cpy)
		return
	}

	for _, v := range graph[i] {
		dps(graph, v, n, res, curPath)
	}
}
