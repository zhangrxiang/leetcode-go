package problemset

import (
	"fmt"
	"testing"
)

//给你一个数组 nums和一个值 val，你需要 原地 移除所有数值等于val的元素，并返回移除后数组的新长度。
//不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
//元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
func removeElement(nums []int, val int) int {
	current := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[current] = nums[i]
			current++
		}
	}
	return current
}

func TestRemoveElement(t *testing.T) {
	data := map[int][]int{
		0: {},
		1: {1, 1, 1, 1, 1},
		2: {1, 1, 1, 1, 1, 2, 2, 2},
		3: {0, 1, 1, 1, 1, 1, 2, 2, 2},
		4: {0, 0, 1, 2, 2, 3, 4, 4, 4, 5},
	}
	res := []int{0, 0, 5, 9, 7}
	for k, v := range data {
		l := removeElement(v, k)
		if res[k] != l {
			t.Fatal(fmt.Sprintf("%d => want %d but %d ,arr %d", k, res[k], l, len(data[k])))
		} else {
			t.Log(fmt.Sprintf("%d => want %d so %d ,arr %d", k, res[k], l, len(data[k])))
		}
	}
}

//执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗：2.1 MB, 在所有 Go 提交中击败了99.92%的用户
