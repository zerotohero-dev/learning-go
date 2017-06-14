package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// DFS
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := root.Left
	right := root.Right

	root.Left = invertTree(right)
	root.Right = invertTree(left)

	return root
}

// BFS
func invertTreeBFS(root *TreeNode) *TreeNode {
	queue := make([]*TreeNode, 0)

	queue = append(queue, root)

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		left := node.Left
		right := node.Right

		node.Left = right
		node.Right = left

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return root
}

func main() {
	tree := TreeNode{42, nil, nil}

	fmt.Println(
		invertTree(&tree),
	)

	fmt.Println(
		invertTreeBFS(&tree),
	)
}

