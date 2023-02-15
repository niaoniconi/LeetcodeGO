package numberTheory

import (
	"sort"
)

// IsGoodArray  是否好数列
/**
其实我已经想到是找唯一共同的非1约数了，但是没想到直接用哦那个最大公约数算法，还以为最大公约数会很难
 */
func IsGoodArray(nums []int)bool{
	sort.Ints(nums)
	if nums[0]==1{
		return true
	}
	divisors:= FindDivisors(nums[0])
	divisorsMap:=make(map[int]int)
	for _,v:=range divisors{
		divisorsMap[v]=0
	}
	for i:=1;i<len(nums);i++{
		if len(divisorsMap)==0{
			return true
		}
		for j,_:=range divisorsMap{
			if nums[i]%j!=0{
				delete(divisorsMap,j)
			}
		}
	}
	//直到最后一个数遍历完，依然有公约数，才能return false
	if len(divisorsMap)==0{
		return true
	}else{
		return false
	}
}

// FindDivisors 返回num除1以外的约数
func FindDivisors(num int) []int{
	var divisors []int
	if num==1 {
		return divisors
	}
	for i:=2;i<=num;i++{
		if num%i==0{
			divisors=append(divisors,i)
		}
		for num%i==0{
			num/=i
		}
	}
	return divisors
}
