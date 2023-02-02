package directedGraph
func ShortestAlternatingPaths(n int,redEdges [][]int,blueEdges [][]int)[]int{  //good news : the lines don't have weight
	answer:=make([]int,n,n)
	mapRB:= make([][]int,n, n)
	for i:=0;i<n;i++ {
		mapRB[i] =make([]int,n,n)
	}
	for _,v :=range redEdges{		//红线是1
		mapRB[v[0]][v[1]]+=1
	}
	for _,v :=range blueEdges{		//蓝线是2，红蓝都有是3
		mapRB[v[0]][v[1]]+=2
	}
	redBegin:=make([]int,n,n)
	redBegin:=make([]int,n,n)
	var findBlueStack,findRedStack []int
	findRedStack=append(findRedStack,0)
	findBlueStack=append(findBlueStack,0)
	for len(findBlueStack)==0||len(findRedStack)==0{
		for _,v:=range findRedStack{


		}
		for _,v:=range findBlueStack{

		}

	}





	return answer
}