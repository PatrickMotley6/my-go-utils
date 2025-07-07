// package stringutil 提供了一系列字符串处理的实用函数。
package stringutil

import "strings"

// Reverse 返回一个将输入字符串 s 按 rune（Unicode码点）反转后的新字符串。
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// IsPalindrome 判断一个字符串 s 是否是回文。
// 这个实现通过调用 Reverse 函数来完成，可读性好。
func IsPalindrome(s string) bool {
	return s == Reverse(s)
}

func ToUpper(s string) string {
	strings.ToUpper(s)
}
