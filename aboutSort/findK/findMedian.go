package findK

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m == 1 && n == 1 {
		return float64((nums2[0] + nums1[0]) / 2)
	}
	if nums1[m-1] <= nums2[0] {
		nums1 = append(nums1, nums2...)
		return float64(nums1[(m+n-1)/2]+nums1[(m+n)/2]) / 2
	}
	if nums2[n-1] <= nums1[0] {
		nums2 = append(nums2, nums1...)
		return float64(nums2[(m+n-1)/2]+nums2[(m+n)/2]) / 2
	}

	if nums1[m/2] <= nums2[n/2] {
		return FindMedianSortedArrays(nums1[m/2:], nums2[:n/2])
	} else {
		return FindMedianSortedArrays(nums1[:m/2], nums2[n/2:])
	}
}
