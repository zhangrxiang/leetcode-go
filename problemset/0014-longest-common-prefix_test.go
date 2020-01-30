package problemset

import (
	"log"
	"testing"
)

//编写一个函数来查找字符串数组中的最长公共前缀。
//如果不存在公共前缀，返回空字符串 ""。
//示例 1:
//输入: ["flower","flow","flight"]
//输出: "fl"
//示例 2:
//输入: ["dog","racecar","car"]
//输出: ""
//解释: 输入不存在公共前缀。
//说明:
//所有输入只包含小写字母 a-z 。
func TestLongestCommonPrefix(t *testing.T) {
	log.Println(longestCommonPrefix([]string{"abcd", "abd", "abd"}))
	log.Println(longestCommonPrefix([]string{"flower", "flow", "flight", "flight"}))
	log.Println(longestCommonPrefix([]string{"abcd", "abcd", "abc"}))
	log.Println(longestCommonPrefix([]string{"11", "22", "33"}))
	log.Println(longestCommonPrefix([]string{"aca", "cba"}))
	log.Println(longestCommonPrefix([]string{"aca"}))
	log.Println(longestCommonPrefix([]string{"cc", "cc"}))
	log.Println(longestCommonPrefix([]string{"cc", "c"}))
	log.Println(longestCommonPrefix([]string{"c", "c"}))
	log.Println(longestCommonPrefix([]string{"c", "b"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	min := len(strs[0])
	index := 0
	for k, v := range strs {
		l := len(v)
		if min > l {
			min = l
			index = k
		}
	}
	for i := 0; i < min; i++ {
		for j := 0; j < len(strs)-1; j++ {
			if strs[j][i] != strs[j+1][i] {
				return strs[j][0:i]
			}
		}
	}
	return strs[index]
}
