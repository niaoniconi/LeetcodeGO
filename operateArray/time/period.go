package time

import (
	"sort"
)

type BriefTime struct{
	Hour int
	Minute int
}

// AlertName https://leetcode.cn/problems/alert-using-same-key-card-three-or-more-times-in-a-one-hour-period/
/**
官方解很多时候就是要精妙很多，你不得不承认
官方思路和我差不多 哈希表+排序
有两个地方需要学习
1.直接计算时间，没有用额外的结构体，
hour := int(t[0]-'0')*10 + int(t[1]-'0')
minute := int(t[3]-'0')*10 + int(t[4]-'0')
这个我记得我吃了一个string的介绍，难道是因为这次都在ascii码内，所以直接这样用没问题吗？可以再测试一下
2.计算超过三次的人时候，直接看当前时间time[i]和time[i+2]是不是在60分钟内
因为已经排好序了，所以只要i+2在，i+1也会在
我的解：
时间复杂度：O(nlogn)
空间复杂度：O(n)，哈希表和答案表
 */
func AlertName(keyName []string,keyTime []string) []string{
	var answer []string
	nameTimes:=make(map[string][]string)
	for i,v:=range keyName{   //O(n)
		if _,ok:=nameTimes[v];ok{
			nameTimes[v]=append(nameTimes[v],keyTime[i])
		}else{
			nameTimes[v]=[]string{keyTime[i]}
		}
	}
	//计算超过三次的人
	for name,times:=range nameTimes{ //O(n) 如果循环是最坏，那么排序就会是最好
		sort.Strings(times) //O(nlogn)
		if len(times)==0 {
			continue
		}
		count:=1
		firstTime:=times[0]
		secondTime:=""
		for i:=1;i<len(times);i++{
			if IsBetweenOneHour(firstTime,times[i]){
				count++
				if count>2{
					secondTime=""
					break
				}
				secondTime=times[i]
			}else{
				if secondTime!=""&&IsBetweenOneHour(secondTime,times[i]){
					firstTime=secondTime
					secondTime=times[i]
					count=2
				}else{
					firstTime=times[i]
					count=1
				}

			}

		}
		if count>2{
			answer=append(answer,name)
		}
	}
	sort.Strings(answer)//O(nlogn)
	return answer
}

// IsBetweenOneHour "10:40","11:00"  time format
/**
return if time2-time1>60 m
 */
func IsBetweenOneHour(time1 string,time2 string) bool{
	time11:=[]rune(time1)
	time22:=[]rune(time2)
	firstTime:=BriefTime{
		Hour: int((time11[0]-'0')*10 + time11[1] - '0'),
		Minute:int((time11[3]-'0')*10 + time11[4] - '0'),
	}
	secondTime:=BriefTime{
		Hour: int((time22[0]-'0')*10 + time22[1] - '0'),
		Minute:int((time22[3]-'0')*10 + time22[4] - '0'),
	}
	diffMinutes:=(secondTime.Hour-firstTime.Hour)*60+secondTime.Minute-firstTime.Minute
	if diffMinutes<=60{
		return true
	}else{
		return false
	}
}
