package main

/*
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。
进阶:

你能不将整数转为字符串来解决这个问题吗？
*/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	y := 0
	num := x
	for x != 0 {
		y = y*10 + x%10
		x = x / 10
	}
	if y != num {
		return false
	}

	return true
}

func main() {
	isPalindrome(-123)
}
