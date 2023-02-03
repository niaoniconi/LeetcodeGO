package main

import (
	"fmt"
	"github.com/LeetcodeGO/operateTree/binaryTree"
)

func main() {

	fmt.Println("hello world")
	node2 := &binaryTree.TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	node3 := &binaryTree.TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	node1 := &binaryTree.TreeNode{
		Val:   1,
		Left:  node2,
		Right: node3,
	}
	answer := binaryTree.BtreeGameWiningMove(node1, 3, 1)
	fmt.Println("answer: ", answer)
}
