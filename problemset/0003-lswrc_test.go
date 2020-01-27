package problemset

import (
	"fmt"
	"strings"
	"testing"
)

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//示例 1:
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//

func TestName(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))      //3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))         //1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))        //3
	fmt.Println(lengthOfLongestSubstring("aab"))           //2
	fmt.Println(lengthOfLongestSubstring("dvdf"))          //3
	fmt.Println(lengthOfLongestSubstring("dvdfddvvfvfvf")) //3
	fmt.Println(lengthOfLongestSubstring("1232345"))       //3
}

func lengthOfLongestSubstring(s string) int {
	str := ""
	l := 0
	for _, v := range s {
		i := strings.IndexByte(str, byte(v))
		if i < 0 {
			str += fmt.Sprintf("%c", v)
		} else {
			str = str[i+1:] + fmt.Sprintf("%c", v)
		}
		if len(str) > l {
			l = len(str)
		}
	}
	return l
}
