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
	testGo()
}

func testGo(){
	defer fmt.Println("in main")
	if err := recover(); err != nil {
		fmt.Println(err)
	}

	panic("unknown err")

}

/*1.给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。*/
func test1(s string) bool{
	stack:=make([]rune,0)
	for _,v:=range s{
		switch v{
		case '(':
			stack=append(stack,v)
		case '{':
			stack=append(stack,v)
		case '[':
			stack=append(stack,v)
		case ')':
			if len(stack)!=0&&stack[len(stack)-1]=='('{
				stack=stack[0:len(stack)-1]
			}else{
				return false
			}
		case '}':
			if len(stack)!=0&&stack[len(stack)-1]=='{'{
				stack=stack[0:len(stack)-1]
			}else{
				return false
			}
		case ']':
			if len(stack)!=0&&stack[len(stack)-1]=='['{
				stack=stack[0:len(stack)-1]
			}else{
				return false
			}
		}

	}
	if len(stack)==0{
		return true
	}else{
		return false
	}
}
