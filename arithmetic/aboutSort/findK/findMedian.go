package findK

// FindMedianSortedArrays https://leetcode.cn/problems/median-of-two-sorted-arrays/
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1)<len(nums2){
		temp:=nums1
		nums1=nums2
		nums2=temp
	}
	m := len(nums1)
	n := len(nums2)
	if n==0{
		if m%2==0{
			return float64(nums1[m/2-1]+nums1[m/2])/2
		}else{
			return float64(nums1[m/2])
		}
	}
	for n>1{
		if nums1[m/2]>nums2[n/2] {
			nums1=nums1[:(m-n/2)]
			nums2=nums2[n/2:]
			m-=n/2
			n-=n/2
		} else {
			nums1=nums1[n/2:]
			nums2=nums2[:n-2/n]
			m-=n/2
			n-=n/2
		}
	}
	if n==m{
		return float64(nums1[0]+nums2[0])/2
	}
	if m%2==0{
		if nums2[0]<nums1[m/2]{
			if nums2[0]<nums1[m/2-1]{
				return float64(nums1[m/2-1])
			}else{
				return float64(nums2[0])
			}
		}else{
			return float64(nums1[m/2])
		}
	}else{
		if nums2[0]<nums1[m/2]{
			if nums2[0]<nums1[m/2-1]{
				return float64(nums1[m/2-1]+nums1[m/2])/2
			}else{
				return float64(nums2[0]+nums1[m/2])/2
			}
		}else{
			if nums2[0]<nums1[m/2+1]{
				return float64(nums2[0]+nums1[m/2])/2
			}else{
				return float64(nums1[m/2+1]+nums1[m/2])/2
			}
		}

	}
}
