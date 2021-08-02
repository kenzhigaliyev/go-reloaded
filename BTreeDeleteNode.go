package student

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node == nil {
		return root
	}
	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		if root.Left == nil {
			new := root.Right
			root = nil
			return new
		} else if root.Right == nil {
			new := root.Left
			root = nil
			return new
		} else {
			new := BTreeMin(root.Right)

			root.Data = new.Data
			root.Right = BTreeDeleteNode(root.Right, new)
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
