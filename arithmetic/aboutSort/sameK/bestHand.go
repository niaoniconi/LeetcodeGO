package sameK

import "sort"

// BestHand  https://leetcode.cn/problems/best-poker-hand/
//官方是hash表+计数，确实效率高一点，O(n)，排序O(nlogn)
//但是无所谓，hash用空间更多
func BestHand(ranks []int,suits []byte) string{
	sort.Ints(ranks)
	handType:=4   		//default
	sameCount:=1
	for i:=0;i<len(ranks);i++{
		count:=1
		for i+1<len(ranks)&&ranks[i+1]==ranks[i]{
			count++
			i++
		}
		if count>sameCount{
			sameCount=count
		}
	}
	if sameCount>=3{
		handType=2
	}else if sameCount==2{
		handType=3
	}

	//https://pkg.go.dev/bytes#Count
	/**
	这一段可以被一下替代，这是bytes	包里面的count函数
	if bytes.Count(suits, suits[:1]) == 5 {
	        return "Flush"
	    }
	func Count(s, sep []byte) int
	计算s中含sep的不重叠个数
	Count counts the number of non-overlapping instances of sep in s. If sep is an empty slice, Count returns 1 + the number of UTF-8-encoded code points in s.
	func main() {
		fmt.Println(bytes.Count([]byte("cheese"), []byte("e")))    //3
		fmt.Println(bytes.Count([]byte("five"), []byte(""))) // before & after each rune   //5
	}
	*/
	flag:=true
	firstSuit:=suits[0]
	for i:=1;i<len(suits);i++{
		if firstSuit!=suits[i]{
			flag=false
		}
	}
	if flag{
		handType=1 //check Flush
	}


	switch handType {
	case 1:
		return "Flush"
	case 2:
		return "Three of a Kind"
	case 3:
		return "Pair"
	default:
		return "High Card"
	}

}
