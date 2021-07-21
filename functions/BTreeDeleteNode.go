package functions

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	// if root == nil {
	// 	return nil
	// } else if root.Data > node.Data {
	// 	BTreeDeleteNode(root.Left, node)
	// } else if root.Data < node.Data {
	// 	BTreeDeleteNode(root.Right, node)
	// } else {
	// 	if root.Left == nil {
	// 		return root.Right
	// 	} else if root.Right == nil {
	// 		return root.Left
	// 	}
	// 	result := &TreeNode{}
	// 	result = FindMin(root.Right)
	// 	root.Right = BTreeDeleteNode(root.Right, result)
	// }
	// return root

	if node == nil {
		return root
	}
	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		if root.Left == nil {
			temp := root.Right
			root = nil
			return temp
		} else if root.Right == nil {
			temp := root.Left
			root = nil
			return temp
		} else {
			temp := BTreeMin(root.Right)

			root.Data = temp.Data
			root.Right = BTreeDeleteNode(root.Right, temp)
		}
	}
	return root

}

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return root
	}

	return BTreeMin(root.Left)
}

// func FindMin(root *TreeNode) *TreeNode {
// 	min := root
// 	if min == nil {
// 		return min
// 	}
// 	for min != nil {
// 		min = min.Left
// 	}
// 	return min
// }
