package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	ter := node
	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		ter.Parent.Left = rplc
	} else if node == node.Parent.Right {
		ter.Parent.Right = rplc
	}
	if rplc != nil {
		rplc.Parent = node.Parent
	}
	return root
}
