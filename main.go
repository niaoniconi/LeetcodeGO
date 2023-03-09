package main

import "fmt"

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
	a:=1
	b:=2
	answer:=a+b
	test:=[]int{1,2,3}
	fmt.Println("answer: ", answer, test[a])
	fmt.Println(ctxStmt,ctxExpr,ctxType,ctxCallee,ctxMultiOK,ctxAssign)
}
