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
