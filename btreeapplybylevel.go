package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	h := BTreeLevelCount(root)
	for i := 0; i < h; i++ {
		applyLevel(root, i, f)
	}
}

func applyLevel(root *TreeNode, level int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	if level == 0 {
		f(root.Data)
	} else if level > 0 {
		applyLevel(root.Left, level-1, f)
		applyLevel(root.Right, level-1, f)
	}
}
