package directedGraph

// ShortestAlternatingPaths https://leetcode.cn/problems/shortest-path-with-alternating-colors/
//注意有环，环可以改变到点的颜色
//到点的颜色不同会影响，点后续走的路
//重点不是一个点一开始是红开还是蓝开，而是这个点后续的状态，我误解了我想要的东西，我把到一个点两种颜色的状态当成了不同的开始
//我一开始甚至有想到环会改变颜色
/**
改善&总结：下次需要把思路和考虑到的点都列出来，把东西写在纸上会好很多，由于在一开始认为边有权重，导致思考的时候添加了很多冗余的东西，尤其是存储
空间：O(n^2+4n)  没用邻接表，用了矩阵
时间：O(r+b+n)  其中 n 是节点数，r 是红色边的数目，b 是蓝色边的数目。广度优先搜索最多访问一个节点两次，最多访问一条边一次，因此时间复杂度为 O(n+r+b)。
*/
func ShortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int { //good news : the lines don't have weight
	answer := make([]int, n, n)
	mapRB := make([][]int, n, n)

	for i := 0; i < n; i++ {
		mapRB[i] = make([]int, n, n)
	}
	//这个懒得优化了，占一点空间怎么了
	for _, v := range redEdges { //红线是1
		mapRB[v[0]][v[1]] += 1
	}
	for _, v := range blueEdges { //蓝线是2，红蓝都有是3
		mapRB[v[0]][v[1]] += 2
	}
	//这个其实也可以优化，但是还是算了，因为我没有设计一个存颜色和节点的对应，诶，好像也可以，用不同的值代表颜色
	redEnd := make([]int, n, n) //到n的颜色是蓝色
	blueEnd := make([]int, n, n)
	var findBlueStack, findRedStack []int
	//这一段是可以优化的，之前是因为同一层级不能乱加而提出来的，现在知道可以暂时搞一个这一层的副本，然后用副本遍历，清空这一层
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

// ShortestAlternatingPathsOfficial 官方解
/**
时间：O(r+b+n)  其中 n 是节点数，r 是红色边的数目，b 是蓝色边的数目。广度优先搜索最多访问一个节点两次，最多访问一条边一次，因此时间复杂度为 O(n+r+b)。
空间复杂度：O(n+r+b)。队列中最多有 2n 个元素，保存next 需要O(r+b) 的空间，保存dist 需要 O(n) 的空间。next指广度优先遍历的队列q
*/
func ShortestAlternatingPathsOfficial(n int, redEdges, blueEdges [][]int) []int {
	//type可以定义在方法里面
	type pair struct{ x, color int }
	g := make([][]pair, n) //颜色表
	for _, e := range redEdges {
		g[e[0]] = append(g[e[0]], pair{e[1], 0})
	}
	for _, e := range blueEdges {
		g[e[0]] = append(g[e[0]], pair{e[1], 1})
	}

	//先遍历 建立可以直接用点的位置做索引找到颜色的 二维数组
	//不过没有的路径的点在这边没有值，这个表更像一个邻接表，而不是矩阵表

	dis := make([]int, n) //答案
	for i := range dis {
		dis[i] = -1
	}
	vis := make([][2]bool, n)    //标明一个点是不是红蓝都有路径
	vis[0] = [2]bool{true, true} //0不需要再找路径了
	q := []pair{{0, 0}, {0, 1}}  //用q来充当广度优点遍历中的栈，一开始有两个起点，分别是0红和0蓝
	for level := 0; len(q) > 0; level++ {
		tmp := q //因为是层次遍历，每一层都要单独存储，不能出一个进一个
		q = nil
		for _, p := range tmp {
			x := p.x
			//初始化，给0节点长度赋0值，也方便后面长度的累加；这一层非0的节点赋值，最先走到的节点一定值比较大
			//我没这样想，是因为一开始觉得每一个边都是有值的，导致后面写题定的时候也没有摆脱这一点
			if dis[x] < 0 {
				dis[x] = level
			}
			for _, to := range g[x] { //看x节点有几条边，能到那个点
				if to.color != p.color && !vis[to.x][to.color] { //只有颜色不一样的，且之前没有才能被选中
					vis[to.x][to.color] = true //命中之后，颜色设置为true
					q = append(q, to)          //添加下一层的节点
				}
			}
		}
	}
	return dis
}

// ShortestAlternatingPathsOptimize
/**
感想1：当你写出一个傻逼数据结构的时候，你就要选择和这个数据结构适配的解法，要不然换个解法的时候，你会发现你的傻逼数据结构会让你浪费很多时间
感想2：当你是要分层的时候，一定要严格记住是在哪一层，官方是把上一层的颜色也存到队列里的，
如果你要单独用一个结构来存当前的color和上一层的color，你就要深拷贝
gain：slice是有深拷贝和浅拷贝的，因为他是一个reference type，深拷贝用copy

用数字存不同颜色太痛苦了，还是应该贴近数据的原始涵义
*/
func ShortestAlternatingPathsOptimize(n int, redEdges, blueEdges [][]int) []int {
	answer := make([]int, n)
	for i := range answer { //赋一个初始值，-1很方便计算
		answer[i] = -1
	}
	mapRB := make([][]int, n, n)
	for i := 0; i < n; i++ {
		mapRB[i] = make([]int, n, n)
	}
	//这个懒得优化了，占一点空间怎么了
	for _, v := range redEdges { //红线是1
		mapRB[v[0]][v[1]] += 1
	}
	for _, v := range blueEdges { //蓝线是2，红蓝都有是3
		mapRB[v[0]][v[1]] += 2
	}
	colorEnd := make([]int, n, n) //红色是1，蓝色是2，红蓝都有是3
	tempEnd := make([]int, n, n)
	//var findBlueStack, findRedStack []int
	//这一段是可以优化的，之前是因为同一层级不能乱加而提出来的，现在知道可以暂时搞一个这一层的副本，然后用副本遍历，清空这一层
	var queue []int
	queue = append(queue, 0)
	colorEnd[0] = 3
	for level := 0; len(queue) != 0; level++ { //学习一下这种优秀的手法，判定条件和level没有关系，但是level可以作为一个计数器
		temp := queue
		queue = nil
		copy(tempEnd, colorEnd)
		//找queue里节点的下一条边就对了，找到了再来校验蓝红
		for _, v := range temp { //queue中有这个节点的时候，就说明有一条到这个节点的路，精妙，之前没注意到这个冗余信息
			//路径和长度是分开的，这道题并不要求找到具体的路径，或者这个路径的重点是红还是蓝，所以很多关于路径的信息都可以忽略
			if answer[v] < 0 {
				answer[v] = level
			}
			for i := 0; i < n; i++ {
				//在这边走路径的时候，也只在意这个点下一条边能走的颜色，其他的都不在意，也不需要记录
				//需要注意有环，一个颜色走一遍就不能再走了
				//注意点
				//mapRB[v][i]和colorEnd[i]的颜色不一样，i节点这个颜色必须没有走过，mapRB[i]必须有值且和colorEnd[v]颜色不一样
				//这个有一个问题，会少算层数，用层数，不存路径长度的坏处就是这样，必须控制层数，严格分层
				if colorEnd[i] == 3 || colorEnd[i] == mapRB[v][i] {
					continue
				}
				if mapRB[v][i] == 3 && colorEnd[i] != 3-tempEnd[v] {
					if tempEnd[v] == 3 {
						colorEnd[i] = 3
					} else {
						colorEnd[i] += 3 - tempEnd[v]
					}
					queue = append(queue, i)
				} else if mapRB[v][i] != 0 {
					if tempEnd[v] == 3 {
						colorEnd[i] += mapRB[v][i]
						queue = append(queue, i)
					} else if tempEnd[v] == 3-mapRB[v][i] {
						colorEnd[i] += mapRB[v][i]
						queue = append(queue, i)
					}

				}
			}
		}

	}
	return answer

}
