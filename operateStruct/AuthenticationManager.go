package operateStruct

// AuthenticationManager token调度器
/**
官方还有一个哈希表+双向链表的方法
链表维护token，并且严格按token的过期时间排列（因为每次调用函数，传入的currentTime都是递增的
然后renew的时候移动节点，
然后在count的时候从头删除，同时删除map的节点，然后map的长度就是剩余节点数
感觉不太有必要的样子
*/
type AuthenticationManager struct {
	TimeToLive int
	Tokens map[string]int      //存储token对应的过期时间
}
func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		TimeToLive:timeToLive,
		Tokens:make(map[string]int),
	}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int)  {  //每个id不会重复，就不叫校验了
	this.Tokens[tokenId]=currentTime+this.TimeToLive
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int)  {
	if expireTime,ok:=this.Tokens[tokenId];ok{
		if currentTime<expireTime{
			this.Tokens[tokenId]=currentTime+this.TimeToLive
		}else{
			delete(this.Tokens,tokenId)
		}
	}
}


func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	count:=0
	for _,expireTime:=range this.Tokens{
		if currentTime<expireTime{
			count++
		}
	}

	return count

}