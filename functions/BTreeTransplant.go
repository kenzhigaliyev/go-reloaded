package functions

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Data > node.Data {
		if root.Left.Data == node.Data {
			root.Left = rplc
			rplc.Parent = root
		} else {
			BTreeTransplant(root.Left, node, rplc)
		}
	} else if root.Data < node.Data {
		if root.Right.Data == node.Data {
			root.Right = rplc
			rplc.Parent = root
		} else {
			BTreeTransplant(root.Right, node, rplc)
		}
	} else {
		return rplc
	}
	return root
}
