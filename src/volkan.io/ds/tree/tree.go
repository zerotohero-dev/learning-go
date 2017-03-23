package tree

// BTree is a binary tree.
type BTree struct {
	Left  *BTree
	Value interface{}
	Right *BTree
}


