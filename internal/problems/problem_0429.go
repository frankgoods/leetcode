package problems

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Q []*Node

func NewQ() *Q {
	return (*Q)(&[]*Node{})
}

func (q *Q) Push(n *Node) {
	*q = append(*q, n)
}

func (q *Q) Pop() *Node {
	v := (*q)[0]
	*q = (*q)[1:len(*q)]
	return v
}

func (q *Q) Len() int {
	return len(*q)
}

func levelOrder(root *Node) [][]int {
	res := [][]int{}
	q := NewQ()

	if root == nil {
		return res
	}

	q.Push(nil)
	q.Push(root)

	for q.Len() != 0 {
		v := q.Pop()
		if v == nil {
			if q.Len() != 0 {
				res = append(res, []int{})
				q.Push(nil)
			}
		} else {
			res[len(res)-1] = append(res[len(res)-1], v.Val)
			for _, ch := range v.Children {
				if ch != nil {
					q.Push(ch)
				}
			}
		}
	}

	return res
}
