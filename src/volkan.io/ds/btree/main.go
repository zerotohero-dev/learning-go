package btree

// BTree is a binary tree.
type BTree struct {
	Left  *BTree
	Value interface{}
	Right *BTree
}

func New(left *BTree, value interface{}, right *BTree) *BTree {
	return &BTree{left, value, right}
}
