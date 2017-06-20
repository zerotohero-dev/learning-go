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
	Val int
	Left *TreeNode
	Right *TreeNode
}

func sumOfLeftLeavesRecursive(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := 0

	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			result += root.Left.Val
		} else {
			result += sumOfLeftLeavesRecursive(root.Left)
		}
	}

	result += sumOfLeftLeavesRecursive(root.Right)

	return result
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := 0

	stack := make([]*TreeNode, 0)

	stack = append(stack, root)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[0:len(stack)-1]

		if node.Left != nil {
			if node.Left.Left == nil && node.Left.Right == nil {
				result += node.Left.Val
			} else {
				stack = append(stack, node.Left)
			}
		}

		if node.Right != nil {
			if node.Right.Left != nil || node.Right.Right != nil {
				stack = append(stack, node.Right)
			}
		}
	}

	return result
}

func main() {
	sumOfLeftLeavesRecursive(nil)
	sumOfLeftLeaves(nil)
}
