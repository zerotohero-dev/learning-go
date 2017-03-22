package ds

// Tree is a binary tree.
type BTree struct {
	Left  *BTree
	Value interface{}
	Right *BTree
}


func prettyFormat(node BTree*, indent string, last bool) string {
	result := indent

	if last {
		result += "\\-"
		indent += "| "
	} else {
		result += "|-"
		indent += "| "
	}

	result += (*node).Value + "\n"

	children := make([]*BTree, 0)

	if node.Left != nil {
		children = append(children, node.Left)
	}

	if node.Right != nil {
		children = append(children, node.Right)
	}

	for i := 0; i < len(children); i++ {
		result += prettyFormat(children[i], indent, i == len(children) - 1)
	}

	return result
}

func (t *BTree) String() string {
	return prettyFormat(t, "", false)
}
