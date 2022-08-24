package problems

import "fmt"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func fillNext(node *Node, myNext *Node) {
	if node == nil {
		return
	}

	node.Next = myNext

	rightNext := myNext
	for rightNext != nil {
		if rightNext.Left != nil {
			rightNext = rightNext.Left
			break
		}
		if rightNext.Right != nil {
			rightNext = rightNext.Right
			break
		}
		rightNext = rightNext.Next
	}

	v := -1
	if myNext != nil {
		v = myNext.Val
	}
	fmt.Println(node.Val, v)

	if node.Right != nil {
		fillNext(node.Right, rightNext)
		fillNext(node.Left, node.Right)
	} else {
		fillNext(node.Left, rightNext)
	}
}

func connect(root *Node) *Node {
	fillNext(root, nil)
	return root
}
