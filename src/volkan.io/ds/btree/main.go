package btree

// BTree is a binary tree.
type BTree struct {
	Left  *BTree
	Value interface{}
	Right *BTree
}


func New(left *BTree, value interface{}, right *BTree) *BTree {
	return &BTree{ left, value, right }
}

/*
	Output:

	1
	1
	2
	3
	5
	8
	13
	21
	34
	55
	89
	144
	233
*/