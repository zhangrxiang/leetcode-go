package problemset

import (
	"log"
	"math"
	"testing"
)

//整数反转
//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
//示例 1:
//输入: 123
//输出: 321
// 示例 2:
//输入: -123
//输出: -321
//示例 3:
//输入: 120
//输出: 21

func TestReverse(t *testing.T) {
	log.Println(reverse(12))
	log.Println(reverse(123))
	log.Println(reverse(-123))
	log.Println(reverse(120))

}
func reverse(x int) int {
	sum := 0
	ten := 1
	if x < 0 {
		ten = -1
		x = -x
	}
	for {
		a := x / 10
		if x > 0 {
			sum = sum*10 + (x % 10)
			x = a
		} else {
			sum = sum * ten
			if sum < math.MinInt32 || sum > math.MaxInt32 {
				return 0
			}
			return sum
		}
	}
}
