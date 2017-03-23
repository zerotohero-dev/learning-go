package tree

import (
	"fmt"
)

// Creates a directory-tree-like representation of the tree
// and returns it as a string.
func prettyFormat(node *BTree, indent string, last bool) string {
	result := indent

	if last {
		result += "\\-"
		indent += "  "
	} else {
		result += "|-"
		indent += "| "
	}

	if node == nil {
		result += "<nada>\n"

		return result
	}

	result += fmt.Sprintf("%v", node.Value) + "\n"

	children := make([]*BTree, 0)

	if node.Left != nil {
		children = append(children, node.Left)
	} else {
		children = append(children, nil)
	}

	if node.Right != nil {
		children = append(children, node.Right)
	} else {
		children = append(children, nil)
	}

	for i := 0; i < len(children); i++ {
		result += prettyFormat(children[i], indent, i == len(children)-1)
	}

	return result
}

func (t *BTree) String() string {
	return prettyFormat(t, "", true)
}
