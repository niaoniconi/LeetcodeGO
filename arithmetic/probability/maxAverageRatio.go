package probability

import "sort"

// MaxAverageRatio  官方想法和我差不多，维护一个队列，以value值的大小为条件，但是这个队列在go里面只能用array了
/**
https://leetcode.cn/problems/longest-well-performing-interval/solution/biao-xian-liang-hao-de-zui-chang-shi-jia-rlij/
 */
func MaxAverageRatio(classes [][]int, extraStudents int)float64{
	var ratio float64
	sort.Slice(classes, func(i, j int) bool {
		return classes[i][1]<classes[j][1]
	})
	for _,value:=range classes{
		if value[0]<value[1]&&extraStudents>0{
			diff:=value[1]-value[0]
			if extraStudents>diff{
				extraStudents-=diff
				value[0]+=diff
			}else{
				value[0]+=extraStudents
				extraStudents=0
			}
		}
		ratio+=float64(value[0])/float64(value[1])
	}
	ratio=ratio/float64(len(classes))
	return ratio
}
