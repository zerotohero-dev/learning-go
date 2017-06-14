package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	val := 0
	if t1 != nil {
		val += t1.Val
	}
	if t2 != nil {
		val += t2.Val
	}
	newNode := TreeNode{val, nil, nil}

	if t1 == nil {
		newNode.Left = mergeTrees(nil, t2.Left)
		newNode.Right = mergeTrees(nil, t2.Right)
	} else if t2 == nil {
		newNode.Left = mergeTrees(t1.Left, nil)
		newNode.Right = mergeTrees(t1.Right, nil)
	} else {
		newNode.Left = mergeTrees(t1.Left, t2.Left)
		newNode.Right = mergeTrees(t1.Right, t2.Right)
	}

	return &newNode
}

func main() {
	t1 := TreeNode{31, nil, nil}
	t2 := TreeNode{44, nil, nil}

	fmt.Println(
		mergeTrees(&t1, &t2),
	)
}
