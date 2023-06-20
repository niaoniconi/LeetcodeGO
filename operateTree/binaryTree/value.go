package binaryTree

// EvaluateTree  https://leetcode.cn/problems/evaluate-boolean-binary-tree/
/**
for full binary Tree
和官方解思路基本一样，只不过官方判断叶节点使用是否有左子结点，我之前也是这样想的，后来还是换成了值，结果都一样
*/
func EvaluateTree(root *TreeNode) bool {
	if root.Val < 2 {
		return root.Val == 1
	} else if root.Val == 2 {
		return EvaluateTree(root.Left) || EvaluateTree(root.Right)
	} else {
		return EvaluateTree(root.Left) && EvaluateTree(root.Right)
	}
}
