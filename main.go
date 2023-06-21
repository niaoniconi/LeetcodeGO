package main

import "github.com/LeetcodeGO/operateTree/binaryTree"

//iota 一款迭代器
const (
	ctxStmt    = 1 << iota // evaluated at statement level
	ctxExpr                // evaluated in value context
	ctxType                // evaluated in type context  4
	ctxCallee              // call-only expressions are ok
	ctxMultiOK             // multivalue function returns are ok
	ctxAssign              // assigning to expression
)

func main() {
	testGo()
}

func testGo() {
	binaryTree.NumTrees(4)
}
