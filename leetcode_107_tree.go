package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue []*TreeNode

func (q *Queue) In(v *TreeNode) {
	*q = append(*q, v)
}

func (q *Queue) Out() *TreeNode {
	if q.IsNil() {
		return nil
	}
	ret := (*q)[0]
	*q = (*q)[1:len(*q)]
	return ret
}

func (q *Queue) IsNil() bool {
	if len(*q) > 0 {
		return false
	} else {
		return true
	}
}

func (q *Queue) Size() int {
	return len(*q)
}

var q Queue

func levelOrderBottom(root *TreeNode) [][]int {
	var ret [][]int

	if root == nil {
		return nil
	}

	q.In(root)
	for i := 0; !q.IsNil(); i++ {
		var layer []int
		for size := q.Size(); size > 0; size-- {
			node := q.Out()
			layer = append([]int{node.Val}, layer...)
			if node.Right != nil {
				q.In(node.Right)
			}
			if node.Left != nil {
				q.In(node.Left)
			}
		}
		ret = append([][]int{layer}, ret...)
	}

	return ret
}

func main() {
}
