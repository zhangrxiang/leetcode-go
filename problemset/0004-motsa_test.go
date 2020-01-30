package problemset

import (
	"log"
	"sort"
	"testing"
)

//median-of-two-sorted-arrays

//给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
//请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
//你可以假设 nums1 和 nums2 不会同时为空。
//示例 1:
//nums1 = [1, 3]
//nums2 = [2]
//则中位数是 2.0
//示例 2:
//nums1 = [1, 2]
//nums2 = [3, 4]
//则中位数是 (2 + 3)/2 = 2.5

func TestFindMedianSortedArrays(t *testing.T) {
	log.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	log.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4, 5}))
	log.Println(findMedianSortedArrays([]int{1, 2, 3}, []int{3, 4, 5}))
	log.Println(findMedianSortedArrays([]int{1, 2, 3, 4}, []int{3, 4, 5}))
}

//执行用时 :28 ms, 在所有 Go 提交中击败了11.82%的用户
//内存消耗 :6 MB, 在所有 Go 提交中击败了23.64%的用户
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	if len(nums)%2 == 0 {
		return float64(nums[len(nums)/2-1]+nums[len(nums)/2]) / 2
	} else {
		return float64(nums[len(nums)/2])
	}
}
