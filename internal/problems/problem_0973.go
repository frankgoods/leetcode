package problems

import "sort"

type Point struct {
	X int
	Y int
	D int
}

func kClosest(points [][]int, k int) [][]int {
	n := len(points)
	closest := make([]*Point, 0, n)

	for _, p := range points {
		closest = append(closest, &Point{p[0], p[1], p[0]*p[0] + p[1]*p[1]})
	}

	sort.Slice(closest, func(i, j int) bool { return closest[i].D < closest[j].D })

	res := make([][]int, 0, k)

	for i := 0; i < k; i++ {
		res = append(res, []int{closest[i].X, closest[i].Y})
	}
	return res
}
