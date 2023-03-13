package simulate

// MinNumberOfHours https://leetcode.cn/problems/minimum-hours-of-training-to-win-a-competition/
/**
因为写错了参数名称有一次错误提交，我的母语是无语，下次要更加注意
 */
func MinNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	hours:=0
	for i,v:=range energy{
		if v>=initialEnergy{
			hours+=v-initialEnergy+1
			initialEnergy=v+1
		}
		if experience[i]>=initialExperience{
			hours+=experience[i]-initialExperience+1
			initialExperience=experience[i]+1
		}
		initialEnergy-=v
		initialExperience+=experience[i]
	}

	return hours
}
