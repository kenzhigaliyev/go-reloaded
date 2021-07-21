package functions

import (
	"fmt"
)

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
	} else if data > root.Data {
		root.Right = BTreeInsertData(root.Right, data)
	}
	return root
}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root.Left != nil {
		BTreeApplyInorder(root.Left, f)
	}
	f(root.Data)
	if root.Right != nil {
		BTreeApplyInorder(root.Right, f)
	}
}

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root.Left != nil {
		BTreeApplyPostorder(root.Left, f)
	}
	if root.Right != nil {
		BTreeApplyPostorder(root.Right, f)
	}
	f(root.Data)
}

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	f(root.Data)
	if root.Left != nil {
		BTreeApplyPreorder(root.Left, f)
	}
	if root.Right != nil {
		BTreeApplyPreorder(root.Right, f)
	}
}

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return root
	}
	if root.Data == elem {
		return root
	}

	if root.Data > elem {
		return BTreeSearchItem(root.Left, elem)
	}
	return BTreeSearchItem(root.Right, elem)
}

func PrintTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Data)
	PrintTree(root.Left)
	PrintTree(root.Right)
}
