package problemset

import (
	"fmt"
	"testing"
)

//给定一个 haystack 字符串和一个 needle 字符串，
//在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
func strStr(haystack string, needle string) int {
	l1 := len(haystack)
	l2 := len(needle)
	if l2 == 0 {
		return 0
	}
	i, j := 0, 0
	for i < l1 && j < l2 {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}
	if l2 == j {
		return i - j
	}
	return -1
}

func TestStrStr(t *testing.T) {
	data := [][]interface{}{
		{"hello", "ll", 2},
		{"abcd", "a", 0},
		{"abcd", "cd", 2},
		{"abcd", "", 0},
		{"", "", 0},
		{"mississippi", "issip", 4},
		{"aaaaa", "bba", -1},
		{"aa", "aaa", -1},
		{"mississippi", "issipi", -1},
		{"", "aaa", -1},
	}
	for k, v := range data {
		l := strStr(v[0].(string), v[1].(string))
		if v[2].(int) != l {
			t.Fatal(fmt.Sprintf("%02d. err:     want %d but %d\n", k, v[2].(int), l))
		} else {
			t.Log(fmt.Sprintf("%02d. success:     want %d so %d\n", k, v[2].(int), l))
		}
	}
}

//执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
//内存消耗：2.2 MB, 在所有 Go 提交中击败了27.84% 的用户
