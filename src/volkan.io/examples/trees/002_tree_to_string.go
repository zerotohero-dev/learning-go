package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}

	result := strconv.Itoa(t.Val)
	left := tree2str(t.Left)
	right := tree2str(t.Right)

	if left == "" && right == "" {
		return result
	}

	if left == "" {
		return result + "()" + "(" + right + ")"
	}

	if right == "" {
		return result + "(" + left + ")"
	}

	return result + "(" + left + ")" + "(" + right + ")"
}

func main() {
	node := TreeNode{4, nil, nil}

	fmt.Println(
		"result:" + tree2str(&node),
	)
}