package problems

import "math"

const (
	Empty = math.MaxInt32
)

type Cell struct {
	I int
	J int
}

type Queue []Cell

func NewQueue() *Queue {
	q := make(Queue, 0)
	return &q
}

func (q *Queue) Push(c Cell) {
	*q = append(*q, c)
}

func (q *Queue) Pop() Cell {
	c := (*q)[0]
	*q = (*q)[1:]
	return c
}

func (q *Queue) Empty() bool {
	return len(*q) == 0
}

func bfs(mat [][]int, q *Queue, result [][]int) {
	m := len(mat)
	n := len(mat[0])

	for !q.Empty() {
		cur := q.Pop()

		i := cur.I
		j := cur.J
		curPrice := result[i][j]

		if i-1 >= 0 && curPrice+1 < result[i-1][j] {
			result[i-1][j] = curPrice + 1
			q.Push(Cell{i - 1, j})
		}

		if i+1 < m && curPrice+1 < result[i+1][j] {
			result[i+1][j] = curPrice + 1
			q.Push(Cell{i + 1, j})
		}

		if j-1 >= 0 && curPrice+1 < result[i][j-1] {
			result[i][j-1] = curPrice + 1
			q.Push(Cell{i, j - 1})
		}

		if j+1 < n && curPrice+1 < result[i][j+1] {
			result[i][j+1] = curPrice + 1
			q.Push(Cell{i, j + 1})
		}
	}
}

func updateMatrix(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])

	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = Empty
		}
	}

	queue := NewQueue()
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				result[i][j] = 0
				queue.Push(Cell{i, j})
			}
		}
	}

	bfs(mat, queue, result)

	return result
}
