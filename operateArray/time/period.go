package time

import (
	"sort"
)

type BriefTime struct{
	Hour int
	Minute int
}

func AlertName(keyName []string,keyTime []string) []string{
	var answer []string
	nameTimes:=make(map[string][]string)
	for i,v:=range keyName{
		if _,ok:=nameTimes[v];ok{
			nameTimes[v]=append(nameTimes[v],keyTime[i])
		}else{
			nameTimes[v]=[]string{keyTime[i]}
		}
	}
	for name,times:=range nameTimes{
		sort.Strings(times)
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
	sort.Strings(answer)
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
