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

/*

func (){

	if root ==nil {
		return nil
	}

	if node.Parent == nil {
		root = rplc
	}

	if node == node.Parent.Left {
		node.Parent.Left = rplc
	} else if node == node.Parent.Right {
		node.Parent.Right = rplc
	}
	rplc.Parent = node.Parent
	return root
}

type TreeNode struct {
	Left, Right, Parent *TreeNode
}

*/
