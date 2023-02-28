package sum

import "sort"

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	var answer [][]int
	items1=append(items1,items2...)
	//weights:=make(map[int]int)
	//for _,v:=range items1{
	//	weights[v[0]]+=v[1]
	//}
	//for _,v:=range weights{
	//	weights[v[0]]+=v[1]
	//}
	sort.Slice(items1,func(i,j int)bool{
		return items1[i][0]<items1[j][0]
	})
	tempValue:=-1		//value大于1
	current:=-1
	for _,v:=range items1{
		if v[0]!=tempValue{
			answer=append(answer,[]int{v[0],v[1]})
			tempValue=v[0]
			current++
		}else{
			answer[current][1]+=v[1]
		}
	}
	return answer
}
