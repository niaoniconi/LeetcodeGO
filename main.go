package main

import (
	"fmt"
<<<<<<< HEAD
	"github.com/LeetcodeGO/simulate"
)

func main() {
	fmt.Println("hello world")
	amout := []int{5, 4, 4}
	answer := simulate.FillCups(amout)
	fmt.Println("answer: ", answer)
=======
)

func main() {
	answers:=make([][]int,5)
	for i:=0;i<5;i++{
		answers[i]=make([]int,2)
	}
	fmt.Println("answer: ", answers)
>>>>>>> 6b6fb03c6f1848f23ad1e9dd27c6335401c622bc
}

func test() {
	//node2 := &binaryTree.TreeNode{
	//	Val:   2,
	//	Left:  nil,
	//	Right: nil,
	//}
	//node3 := &binaryTree.TreeNode{
	//	Val:   3,
	//	Left:  nil,
	//	Right: nil,
	//}
	//node1 := &binaryTree.TreeNode{
	//	Val:   1,
	//	Left:  node2,
	//	Right: node3,
	//}
	//answer := binaryTree.BtreeGameWiningMove(node1, 3, 1)
	//fmt.Println("answer: ", answer)
	//red := [][]int{{1, 5}, {2, 2}, {5, 5}, {3, 0}, {4, 5}, {2, 4}, {4, 1}, {1, 0}, {1, 2}, {5, 2}, {2, 3}, {0, 1}}
	//blue := [][]int{{4, 4}, {2, 5}, {1, 1}, {5, 4}, {3, 3}}
	//answer2 := directedGraph.ShortestAlternatingPathsOptimize(6, red, blue)
	//fmt.Println("answer: ", answer2)
	//coins := []int{4, 5, 1, 1, 5, 4, 1, 1, 1, 2}
	//answer := sum.GetMaximumConsecutive(coins)
}
