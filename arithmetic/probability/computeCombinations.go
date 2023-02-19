package probability
//DieSimulator https://leetcode.cn/problems/dice-roll-simulation/
/**
概率论学的一般还总想用概率论解出来，瑞思拜自己
官方是用动态规划写的，先试试
 */
//func DieSimulator(n int,rollMax []int)int{
//	count:=0
//	dp:=make([][]int,n)
//	for i:=0;i<n;i++{
//		dp[i]=make([]int,7)
//	}
//	for i:=1;i<7;i++{
//		dp[]
//	}
//
//
//	return count
//}

func DieSimulatorTryMath(n int,rollMax []int)int{
	count:=0
	//由于rollMax[i]>=1
	mod:=1000000007
	count+=6*5^(n-1)%mod
	//

	return count
}

