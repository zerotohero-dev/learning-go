package main

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode = nil

	for head != nil {
		next := head.Next
		head.Next = newHead
		newHead = head
		head = next
	}

	return newHead
}

func reverseListRec(head, newHead *ListNode) *ListNode {
	if head == nil {
		return newHead
	}

	next := head.Next
	head.Next = newHead

	return reverseListRec(next, head)
}

func reverseListRecursive(head *ListNode) *ListNode {
	return reverseListRec(head, nil)
}

func main() {
	reverseList(nil)
	reverseListRecursive(nil)
}

