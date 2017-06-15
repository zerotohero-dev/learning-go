package main

import "fmt"

const MaxUint = ^uint(0)
//const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func minimum(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func getMinimumDifferenceImpl(min *int, prev *int, root *TreeNode) {

	// The in-order traversal of a binary search tree will be in sorted order.

	if root == nil {
		return
	}

	getMinimumDifferenceImpl(min, prev, root.Left)

	if *prev > MinInt {
		*min = minimum(*min, root.Val - *prev)
	}
	*prev = root.Val

	getMinimumDifferenceImpl(min, prev, root.Right)
}

func getMinimumDifference(root *TreeNode) int {
	min := MaxInt
	prev := MinInt

	getMinimumDifferenceImpl(&min, &prev, root)

	return min
}

func main() {
	fmt.Println(
		getMinimumDifference(nil),
	)
}