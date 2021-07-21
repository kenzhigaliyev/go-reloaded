package functions

import (
	"fmt"
)

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	// if root == nil {
	// 	return root
	// }
	// if root.Data == node.Data {
	// 	root.Data = rplc.Data
	// 	return root
	// }
	// if root.Data > node.Data {
	// 	return BTreeTransplant(root.Left, node, rplc)
	// }
	// return BTreeTransplant(root.Right, node, rplc)

	find := BTreeSearchItem(root, node.Data)
	fmt.Println(find.Data)
	Previous := find
	if find.Parent != nil {
		Previous = find.Parent
	}
	fmt.Println(Previous.Data)
	if Previous.Data < find.Data {
		find.Parent.Right = rplc
	} else {
		find.Parent.Left = rplc
	}
	rplc.Parent = Previous
	return root

}
