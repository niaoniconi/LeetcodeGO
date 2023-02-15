package numberTheory

import (
	"sort"
)

// IsGoodArray  是否好数列
/**
其实我已经想到是找唯一共同的非1约数了，但是没想到直接用哦那个最大公约数算法，还以为最大公约数会很复杂
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

// IsGoodArrayOfficial  官方解：使用gcd
/**
本来我想的是先用前两个数算出一个g，再用g和后面的数一起算，但是这样会很麻烦，需要考虑到的边界条件会很高
直接设定0和任何数公约数是任何数会好很多，虽然也不知道公约数是不是这样规定的，行吧，0没有公约数
其实可以g:=nums[0]，两个相同的数最大公约数是对方，
g==1直接返回true这个也很棒，只要有一个不行就直接退出
 */
func IsGoodArrayOfficial(nums []int)bool{
	g:=0
	for _,value:=range nums{
		g=GCD(g,value)
		if g==1{
			return true
		}
	}
	return false
}

// GCD  0和x的公约数是x，辗转相除法
func GCD(a,b int)int{
	for a!=0{
		a,b=b%a,a
	}
	return b
}