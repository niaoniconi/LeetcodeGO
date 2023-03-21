package main

import (
	"fmt"
	"reflect"
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
type animal interface{
	say(words string)
}
type cat struct {
	language string
}
func (c cat) say(words string){
	println(c.language+" "+words)
}
func main() {
	typeOfError := reflect.TypeOf((*animal)(nil)).Elem()
	customErrorPtr := reflect.TypeOf(&cat{})
	customError := reflect.TypeOf(cat{})

	fmt.Println(customErrorPtr.Implements(typeOfError)) // #=> true
	fmt.Println(customError.Implements(typeOfError)) // #=> false

}
