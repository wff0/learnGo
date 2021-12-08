package offer

// https://leetcode-cn.com/problems/a7VOhD/
func countSubstrings(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < 2*n-1; i++ {
		l, r := i/2, i/2+i%2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			ans++
		}
	}
	return ans
}

/*
主要通过枚举回文中心来解决，要考虑奇数和偶数情况
*/
