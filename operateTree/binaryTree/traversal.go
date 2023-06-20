package binaryTree

import "fmt"

/*
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//InorderTraversal 递归版 https://leetcode.cn/problems/binary-tree-inorder-traversal/submissions/
/**
题解直接用的闭包，确实，go不必非要自己再起一个函数，而且这样也不用考虑传参，
还有一种Morris中序遍历，主要是降运行中内存O(1)，非递归情况
核心操作是增加从左子树到根的回路，让当前节点中左子树的最右节点（左子树中根遍历的最后一个节点），右孩子指向当前节点
但是每个节点会被访问两次，O(2n)复杂度
这样每个节点都要两个操作，判断是否链接，未链接：链接，继续往左深入；已链接：打印节点，删除链接（不删行不行，应该不行），走右子树
*/
func InorderTraversal(root *TreeNode) []int {
	var answer []int
	traversal(root, &answer)
	fmt.Println(answer)
	return answer
}
func traversal(root *TreeNode, answer *[]int) {
	if root == nil {
		return
	}
	traversal(root.Left, answer)
	*answer = append(*answer, root.Val)
	traversal(root.Right, answer)
}
