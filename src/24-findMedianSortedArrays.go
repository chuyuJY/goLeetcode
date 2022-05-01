package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 首先，解决奇偶的影响
	len1, len2 := len(nums1), len(nums2)
	midLeft, midRight := (len1+len2+1)/2, (len1+len2+2)/2 // 如果是奇，则midLeft和midRight相等，若为偶，则midRight=midLeft+1
	return 0.5 * float64(getKthSmallest(nums1, 0, len1-1, nums2, 0, len2-1, midLeft)+getKthSmallest(nums1, 0, len1-1, nums2, 0, len2-1, midRight))
}
func getKthSmallest(nums1 []int, start1, end1 int, nums2 []int, start2, end2 int, k int) int {
	// 注意此处的len不是整体数组长度，而是剩余数组长度
	len1, len2 := end1-start1+1, end2-start2+1
	// 此步只是为了保证len1 <= len2, 这样的话，若发生len<k或空的情况，只会发生在len1中，便于后续操作而已
	if len1 > len2 {
		return getKthSmallest(nums2, start2, end2, nums1, start1, end1, k)
	}
	// 若len1空了，只需要返回nums2中从start2开始数的第k个即可
	if len1 == 0 {
		return nums2[start2+k-1]
	}
	// 若”正常“”顺利“执行到最后一步
	if k == 1 {
		return min(nums1[start1], nums2[start2])
	}
	// 这样写，你考虑了start1+len1 < start1+k/2-1 ？？？不可以的！！！
	// i, j := start1+k/2-1, start2+k/2-1
	// 要这样写才对！！！
	i := start1 + min(k/2, len1) - 1
	j := start2 + min(k/2, len2) - 1
	if nums1[i] >= nums2[j] {
		return getKthSmallest(nums1, start1, end1, nums2, j+1, end2, k+start2-1-j) // 注意此处不能是 k-k/2 因为有可能只排除min(start2+k/2-1, len2-1)个
	} else {
		return getKthSmallest(nums1, i+1, end1, nums2, start2, end2, k+start1-1-i)
	}
}

//func min(a, b int) int {
//	if a <= b {
//		return a
//	} else {
//		return b
//	}
//}

func main() {

}
