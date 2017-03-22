package ds

// Tree is a binary tree.
type BTree struct {
	Left  *Tree
	Value interface{}
	Right *Tree
}
