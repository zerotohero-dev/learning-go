package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func max(x int, y int) int {
	if x < y {
		return y
	}

	return x
}


func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1+max(maxDepth(root.Left), maxDepth(root.Right))
}

func main() {
	node := TreeNode{4, nil, nil}

	fmt.Println(
		maxDepth(&node),
	)
}
