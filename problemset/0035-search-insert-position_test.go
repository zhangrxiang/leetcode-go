package problemset

import (
	"fmt"
	"testing"
)

func searchInsert(nums []int, target int) int {
	l := len(nums)
	i := 0
	if target > nums[l-1] {
		return l
	}
	for i < l-1 {
		if nums[i] == target {
			return i
		} else if nums[i] < target && target <= nums[i+1] {
			return i + 1
		}
		i++
	}
	return 0
}

func TestSearchInsert(t *testing.T) {
	data := [][]interface{}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
		{[]int{1}, 0, 0},
		{[]int{1, 3}, 3, 1},
	}
	for k, v := range data {
		l := searchInsert(v[0].([]int), v[1].(int))
		if v[2].(int) != l {
			t.Fatal(fmt.Sprintf("%02d. err:     want %d but %d\n", k, v[2].(int), l))
		} else {
			t.Log(fmt.Sprintf("%02d. success:     want %d so %d\n", k, v[2].(int), l))
		}
	}
}

//执行用时：4 ms, 在所有 Go 提交中击败了90.77% 的用户
//内存消耗：3 MB, 在所有 Go 提交中击败了57.20% 的用户
