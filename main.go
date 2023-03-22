package main

import (
	"fmt"
	"reflect"
)

type node struct {
	children map[rune]*node
	s        rune
	isSet    bool
}

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
func (n *node) insert(word string) {
	i := 0
	cur := n
	temp := n
	for i < len(word) && temp != nil {
		children := temp.children
		if child, ok := children[rune(word[i])]; ok {
			temp = child
			i++
		} else {
			for i < len(word) {
				ch1 := rune(word[i])
				children[ch1] = &node{
					s:        rune(word[i]),
					children: make(map[rune]*node),
				}
				cur = children[ch1]
				if i == len(word)-1 {
					cur.isSet = true
				}
				children = children[rune(word[i])].children
				i++
			}
		}
	}
	temp.isSet = true
}

func (n *node) search(word string) bool {
	temp := n
	for _, v := range word {
		if temp.children[rune(v)] == nil {
			return false
		} else {
			temp = temp.children[rune(v)]
		}
	}
	if temp.isSet {
		return true
	}
	return false
}

func (n *node) startWith(word string) bool {
	for _, v := range word {
		if n.children[rune(v)] == nil {
			return false
		} else {
			n = n.children[rune(v)]
		}
	}
	return true
}

func main() {
	typeOfError := reflect.TypeOf((*animal)(nil)).Elem()
	customErrorPtr := reflect.TypeOf(&cat{})
	customError := reflect.TypeOf(cat{})

	fmt.Println(customErrorPtr.Implements(typeOfError)) // #=> true
	fmt.Println(customError.Implements(typeOfError)) // #=> false

}
