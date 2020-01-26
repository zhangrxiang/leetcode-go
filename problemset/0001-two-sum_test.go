package problemset

import (
	"log"
	"testing"
)

//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
//示例:
//给定 nums = [2, 7, 11, 15], target = 9
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]

func TestTwoSum(t *testing.T) {
	log.Println(twoSum([]int{2, 7, 11, 15}, 9))
	log.Println(twoSum([]int{7, 11, 15, 2}, 9))
	log.Println(twoSum([]int{11, 15, 2, 8}, 9))
	log.Println(twoSum([]int{2, 5, 5, 10}, 10))
	log.Println(twoSum([]int{5, 5, 5, 10}, 10))
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if target-nums[j] == nums[i] {
				return []int{i, j}
			}
		}
	}
	return nil
}
