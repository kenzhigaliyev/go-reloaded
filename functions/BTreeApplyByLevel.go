package functions

func NewLine(root *TreeNode, f func(...interface{}) (int, error), val int) {
	if root == nil {
		return
	}
	if val == 0 {
		f(root.Data)
	}
	if val > 0 {
		NewLine(root.Left, f, val-1)
		NewLine(root.Right, f, val-1)
	}
}

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	counter1 := BTreeLevelCount(root.Left)
	counter2 := BTreeLevelCount(root.Right)
	if counter1 > counter2 {
		return counter1 + 1
	}
	return counter2 + 1
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	levels := BTreeLevelCount(root)
	for i := 0; i < levels; i++ {
		NewLine(root, f, i)
	}
}
