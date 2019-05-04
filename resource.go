package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*----------------------- Stack -----------------------*/
type Stack []*TreeNode

func (s *Stack) Push(v *TreeNode) {
	*s = append(*s, v)
}

func (s *Stack) Pop() *TreeNode {
	if s.IsNil() {
		return nil
	}
	ret := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return ret
}

func (s *Stack) IsNil() bool {
	if len(*s) > 0 {
		return false
	} else {
		return true
	}
}

var s Stack

func testStack() {
	fmt.Printf("----------------------- testStack -----------------------\n")
	s.Push(&TreeNode{Val: 0, Left: nil, Right: nil})
	s.Push(&TreeNode{Val: 1, Left: nil, Right: nil})
	s.Push(&TreeNode{Val: 2, Left: nil, Right: nil})
	s.Push(&TreeNode{Val: 3, Left: nil, Right: nil})
	s.Push(&TreeNode{Val: 4, Left: nil, Right: nil})
	s.Push(&TreeNode{Val: 5, Left: nil, Right: nil})
	s.Push(&TreeNode{Val: 6, Left: nil, Right: nil})
	s.Push(&TreeNode{Val: 7, Left: nil, Right: nil})

	node := s.Pop()
	for node != nil {
		fmt.Printf("%d\n", node.Val)
		node = s.Pop()
	}
}

/*----------------------- Queue -----------------------*/
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

func testQueue() {
	fmt.Printf("----------------------- testQueue -----------------------\n")
	q.In(&TreeNode{Val: 0, Left: nil, Right: nil})
	q.In(&TreeNode{Val: 1, Left: nil, Right: nil})
	q.In(&TreeNode{Val: 2, Left: nil, Right: nil})
	q.In(&TreeNode{Val: 3, Left: nil, Right: nil})
	q.In(&TreeNode{Val: 4, Left: nil, Right: nil})
	q.In(&TreeNode{Val: 5, Left: nil, Right: nil})
	q.In(&TreeNode{Val: 6, Left: nil, Right: nil})
	q.In(&TreeNode{Val: 7, Left: nil, Right: nil})

	node := q.Out()
	for node != nil {
		fmt.Printf("%d\n", node.Val)
		node = q.Out()
	}
}

/*----------------------- xxxxx -----------------------*/

func main() {
	testStack()
	testQueue()
}
