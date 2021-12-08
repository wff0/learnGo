package offer

import (
	"strings"
	"unicode"
)

// https://leetcode-cn.com/problems/XltzEq/
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	n := len(s)
	if n == 0 {
		return true
	}
	l, r := 0, n-1
	for l < r {
		for l < r && !unicode.IsLetter(rune(s[l])) && !unicode.IsDigit(rune(s[l])) {
			l++
		}
		for l < r && !unicode.IsLetter(rune(s[r])) && !unicode.IsDigit(rune(s[r])) {
			r--
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
