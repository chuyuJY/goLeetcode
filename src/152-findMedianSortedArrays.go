package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 消除奇偶影响
	len1, len2 := len(nums1), len(nums2)
	midLeft, midRight := (len1+len2+1)/2, (len1+len2+2)/2 // 如果是奇，则midLeft和midRight相等，若为偶，则midRight=midLeft+1
	return 0.5 * float64((getKthSmallest(nums1, nums2, midLeft) + getKthSmallest(nums1, nums1, midRight)))
}

func getKthSmallest(nums1, nums2 []int, k int) int {
	// 保证只有nums1会出现为空的情况
	if len(nums1) > len(nums2) {
		return getKthSmallest(nums2, nums1, k)
	}
	if len(nums1) == 0 {
		return nums2[k-1]
	}
	if k == 1 {
		return min(nums1[0], nums2[0])
	}
	index := min(k/2, len(nums1))
	if nums1[index-1] < nums2[index-1] {
		return getKthSmallest(nums1[index:], nums2, k-index)
	} else {
		return getKthSmallest(nums1, nums2[index:], k-index)
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
