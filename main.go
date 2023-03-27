package main

import (
	"fmt"
	"github.com/LeetcodeGO/weekContest"
)

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
	testLC()
}

func testLC() {
	//arr := []int{10, 13, 17, 21, 15, 15, 9, 17, 22, 22, 13}
	//answer := subArray.FindLengthOfShortestSubarray(arr)
	//answer := weekContest.KItemsWithMaximumSum(2, 3, 4, 6)
	answer := weekContest.PrimeSubOperation([]int{38, 39, 52, 16, 67, 96, 69, 44, 63, 29})
	fmt.Println(answer)
}
