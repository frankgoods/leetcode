package problems

type Vertex struct {
	X, Y int
}

func dist(grid [][]int, v Vertex) {
	grid[v.X][v.Y] = 100
	q := []Vertex{v}

	for len(q) != 0 {
		cur := q[0]
		curDist := grid[cur.X][cur.Y]
		q = q[1:]
		q = add(grid, q, cur.X+1, cur.Y, curDist+1)
		q = add(grid, q, cur.X-1, cur.Y, curDist+1)
		q = add(grid, q, cur.X, cur.Y+1, curDist+1)
		q = add(grid, q, cur.X, cur.Y-1, curDist+1)
	}
}

func add(grid [][]int, q []Vertex, i, j, curDist int) []Vertex {
	if i < 0 || i == len(grid) || j < 0 || j == len(grid[0]) ||
		(grid[i][j] != 0 && curDist >= grid[i][j]) {
		return q
	}
	grid[i][j] = curDist
	q = append(q, Vertex{i, j})
	return q
}

func maxDistance(grid [][]int) int {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				dist(grid, Vertex{i, j})
			}
		}
	}

	max := -1
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > 100 && (max == -1 || grid[i][j] > max) {
				max = grid[i][j]
			}
		}
	}

	if max > 100 {
		return max - 100
	}
	return max
}
