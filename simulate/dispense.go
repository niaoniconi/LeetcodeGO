package simulate

import "sort"

// FillCups  https://leetcode.cn/problems/minimum-amount-of-time-to-fill-cups/
func FillCups(amount []int) int {
	answer := 0
	sort.Ints(amount)
	max := amount[2]
	mid := amount[1]
	min := amount[0]
	temp := 0
	for min > 0 {
		if max == mid && mid == min {
			answer += max*3/2 + max*3%2
			max = 0
			mid = 0
			min = 0
		} else {
			max--
			mid--
			answer++
		}
		if min > mid {
			if min > max {
				temp = min
				min = mid
				mid = max
				max = temp
			} else {
				temp = min
				min = mid
				mid = temp
			}
		}
	}
	answer += max
	return answer
}

// FillCupsOfficial
/**
官方解：谈心+分类讨论
1.x+y<=z，这个时候最大是z
2.x+y>z，这时候，需要先同等的减少x,y，再尽量让剩下的restX+restY=Z，争取每一次都减少两个容量
t=x+y-z,需要让x，y都同时减少2/t，这样restX+restY=Z
如果t是偶数，answer=t/2+z=(x+y-z)/2+z=(x+y+z)/2=0,     x+y=2n+z=>x+y+z=2n+2z，x+y+z是偶数
如果t是奇数，那么需要考虑，x，y是同时减少t+1/2,还是t-1/2;   x+y=2n+1+z=>x+y+z=2n+2z+1,x+y+z是奇数
t+1/2:restX+restY=Z-1,answer=t+1/2+z=(x+y-z+1)/2+z=(x+y+z+1)/2
t-1/2:restX+restY=Z+1,answer=t-1/2+z+1=(x+y-z-1)/2+z+1=(x+y+z+1)/2
由于/运算符的特性 x+y+z是偶数 (x+y+z+1)/2=(x+y+z)/2
*/
func FillCupsOfficial(amount []int) int {
	sort.Ints(amount)
	if amount[2] >= amount[0]+amount[1] {
		return amount[2]
	} else {
		return (amount[0] + amount[1] + amount[2] + 1) / 2
	}
}
