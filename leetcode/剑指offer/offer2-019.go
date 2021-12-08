package offer

// https://leetcode-cn.com/problems/RQku0D/
func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	valid := func(i, j int) bool {
		left, right := i, j
		for left < right {
			if s[left] == s[right] {
				left++
				right--
			} else {
				return false
			}
		}
		return true
	}
	for l < r {
		if s[l] == s[r] {
			l++
			r--
		} else {
			return valid(l, r-1) || valid(l+1, r)
		}
	}
	return true
}
