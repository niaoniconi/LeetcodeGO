package directedGraph

// ShortestAlternatingPaths https://leetcode.cn/problems/shortest-path-with-alternating-colors/
//注意有环，环可以改变到点的颜色
//到点的颜色不同会影响，点后续走的路
//重点不是一个点一开始是红开还是蓝开，而是这个点后续的状态，我误解了我想要的东西，我把到一个点两种颜色的状态当成了不同的开始
//我一开始甚至有想到环会改变颜色
/**
改善：下次需要把思路和考虑到的点都列出来
*/
func ShortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int { //good news : the lines don't have weight
	answer := make([]int, n, n)
	mapRB := make([][]int, n, n)
	for i := 0; i < n; i++ {
		mapRB[i] = make([]int, n, n)
	}
	for _, v := range redEdges { //红线是1
		mapRB[v[0]][v[1]] += 1
	}
	for _, v := range blueEdges { //蓝线是2，红蓝都有是3
		mapRB[v[0]][v[1]] += 2
	}
	redEnd := make([]int, n, n) //到n的颜色是蓝色
	blueEnd := make([]int, n, n)
	var findBlueStack, findRedStack []int
	//找到最后一条边是红色的
	for i := 1; i < n; i++ {
		if mapRB[0][i] == 1 || mapRB[0][i] == 3 {
			if redEnd[i] == 0 { //只有等于0的会被替代，这是深度优先遍历，line无权重，只要先找到就行
				redEnd[i] = redEnd[0] + 1
				findBlueStack = append(findBlueStack, i)
			}
		}
	}

	//找到最后一条边是蓝色的
	for i := 1; i < n; i++ {
		if mapRB[0][i] == 2 || mapRB[0][i] == 3 {
			if blueEnd[i] == 0 {
				blueEnd[i] = blueEnd[0] + 1
				findRedStack = append(findRedStack, i)
			}
		}
	}

	for len(findBlueStack) != 0 || len(findRedStack) != 0 {
		//红的继续找蓝，蓝的继续找红
		for _, v := range findRedStack {
			if v == 0 {
				continue
			}
			for i := 0; i < n; i++ {
				if mapRB[v][i] == 1 || mapRB[v][i] == 3 {
					if redEnd[i] == 0 && blueEnd[v] != 0 { //只有等于0的会被替代，这是深度优先遍历，line无权重，只要先找到就行
						redEnd[i] = blueEnd[v] + 1
						findBlueStack = append(findBlueStack, i)
					}
				}
			}
		}

		findRedStack = findRedStack[0:0]

		for _, v := range findBlueStack {
			if v == 0 {
				continue
			}
			for i := 0; i < n; i++ {
				if mapRB[v][i] == 2 || mapRB[v][i] == 3 {
					if blueEnd[i] == 0 && redEnd[v] != 0 {
						blueEnd[i] = redEnd[v] + 1
						findRedStack = append(findRedStack, i)
					}
				}
			}
		}
		findBlueStack = findBlueStack[0:0]
	}
	for i := 0; i < n; i++ {
		if redEnd[i] == 0 && blueEnd[i] == 0 {
			answer[i] = -1
		} else if (redEnd[i] != 0 && redEnd[i] < blueEnd[i]) || blueEnd[i] == 0 {
			answer[i] = redEnd[i]
		} else {
			answer[i] = blueEnd[i]
		}
	}
	answer[0] = 0
	return answer
}
