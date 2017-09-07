package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func abs(val int) int {
	if val < 0 {
		return -val
	}

	return val
}

func sum(tilt *int, root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := sum(tilt, root.Left)
	right := sum(tilt, root.Right)

	*tilt += abs(left-right)

	return left + right + root.Val
}

func findTilt(root *TreeNode) int {
	result := 0

	sum(&result, root)

	return result
}

func main() {
	findTilt(nil)
}