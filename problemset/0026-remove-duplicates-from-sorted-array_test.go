package problemset

import (
	"fmt"
	"testing"
)

//26. 删除排序数组中的重复项
//
//给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
//
//不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
//
//
//示例 1:
//
//给定数组 nums = [1,1,2],
//
//函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。
//
//你不需要考虑数组中超出新长度后面的元素。
//
//示例 2:
//
//给定 nums = [0,0,1,1,1,2,2,3,3,4],
//
//函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。
//

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	nums = nums[0:slow]
	return slow + 1
}

func TestRemoveDuplicates(t *testing.T) {
	data := map[int][]int{
		0: {},
		1: {1, 1, 1, 1, 1},
		2: {1, 1, 1, 1, 1, 2, 2, 2},
		3: {0, 1, 1, 1, 1, 1, 2, 2, 2},
		5: {0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
	}
	for k, v := range data {
		l := removeDuplicates(v)
		if k != l {
			t.Fatal(fmt.Sprintf("want %d but %d ,arr %d", k, l, len(data[k])))
		} else {
			t.Log(fmt.Sprintf("want %d but %d ,arr %d", k, l, len(data[k])))
		}
	}
}

func TestSlice(t *testing.T) {
	s := []byte{1, 2, 3, 4, 5, 6}
	fmt.Println(s[0:2], s[2:])
	fmt.Println(s[0:0], s[6:6])
}

//数组完成排序后，我们可以放置两个指针 iii 和 jjj，其中 iii 是慢指针，而 jjj 是快指针。
//只要 nums[i]=nums[j]nums[i] = nums[j]nums[i]=nums[j]，我们就增加 jjj 以跳过重复项。
//当我们遇到 nums[j]≠nums[i]nums[j] \neq nums[i]nums[j]​=nums[i] 时，跳过重复项的运行已经结束，
//因此我们必须把它（nums[j]nums[j]nums[j]）的值复制到 nums[i+1]nums[i + 1]nums[i+1]。
//然后递增 iii，接着我们将再次重复相同的过程，直到 jjj 到达数组的末尾为止。
