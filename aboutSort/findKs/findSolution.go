package findKs

// FindSolution https://leetcode.cn/problems/find-positive-integer-solution-for-a-given-equation/
/*
1<=x,y<=1000
函数是严格增函数,返回正整数对
This is the declaration of customFunction API.
 * @param  x    int
 * @param  x    int
 * @return 	    Returns f(x, y) for any given positive integers x and y.
 *			    Note that f(x, y) is increasing with respect to both x and y.
 *              i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
 */
/**
官方解有三个方法：
1.暴力解，遍历 O(mn)
2.二分查找，对每一个x用查找函数，找到f(x,y)>=z,计算f(x,y)==z   O(mlogn)
使用sort.Search函数
3. 双指针 O(m+n),最好的方法，牢记双指针，注意规律。之前有想过，找到最左边和最右边，但是一直没想过引入边界值，注意，用边界值并不可耻
下面我将用双指针求解
三种方法都有用到1000，不要那么排斥用到边界值，计算机的宿命就不是求出全部解，而是有限解
 */
func FindSolution(customFunction func(int,int)int,z int)[][]int{
	var answer [][]int
	maxj:=1000
	for i:=0;i<=1000;i++{
		for customFunction(i,maxj)>z{
			maxj--
		}
		if maxj<0{
			break
		}
		if customFunction(i,maxj)==z{
			answer=append(answer,[]int{i,maxj})
		}
	}
	return answer
}

// FindSolutionViolent  写暴力解的感觉爽死我了
func FindSolutionViolent(customFunction func(x int,y int)int,z int)[][]int{
	var answer [][]int
	for i:=1;i<=1000;i++{
		for j:=1;j<=1000;j++{
			if customFunction(i,j)==z{
				answer=append(answer,[]int{i,j})
			}
		}
	}
	return answer
}
