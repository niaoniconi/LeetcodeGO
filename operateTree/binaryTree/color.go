package binaryTree

// BtreeGameWiningMove	https://leetcode.cn/problems/binary-tree-coloring-game/
/*
树中节点数目为 n
1 <= x <= n <= 100
n 是奇数
1 <= Node.val <= n
树中所有值 互不相同
*/
func BtreeGameWiningMove(root *TreeNode, n int, x int) bool {
	//x可以染色左子树，右子树，以及其他三部分，只要其中任何一部分大于另两者之和+1，就可以找到y节点，否则找不到
	//官解：换句官方解的话，就是任何一部分大于(n+1)/2，就行，这样这一部分已经大于其他两部分
	//先遍历，找到x，再计算左子树，右子树的节点数
	xNode := FindXInTree(root, x)
	leftNodes := CalculateNodesOfTree(xNode.Left)
	rightNodes := CalculateNodesOfTree(xNode.Right)
	if leftNodes > n-leftNodes || rightNodes > n-rightNodes || n > 2*(leftNodes+rightNodes+1) {
		return true
	}
	return false

}

//官方解：这两个函数合成getSubtreeSize，先用getSubtreeSize函数找到xNode，再用这个函数计算getSubtreeSize函数的左右节点，思路一致
//和官方解的时间复杂度一致O(n)
//见识一下函数变量的骚操作，可以学学
/**
var getSubtreeSize func(*TreeNode) int
    getSubtreeSize = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        if node.Val == x {
            xNode = node
        }
        return 1 + getSubtreeSize(node.Left) + getSubtreeSize(node.Right)
    }
*/

func FindXInTree(root *TreeNode, x int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == x {
		return root
	}
	left := FindXInTree(root.Left, x)
	if left != nil {
		return left
	}
	right := FindXInTree(root.Right, x)
	if right != nil {
		return right
	}
	return nil
}

func CalculateNodesOfTree(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		return 1 + CalculateNodesOfTree(root.Left) + CalculateNodesOfTree(root.Right)
	}
}
