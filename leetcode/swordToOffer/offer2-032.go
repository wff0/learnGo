package offer

// https://leetcode-cn.com/problems/dKk3P7/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) || s == t {
		return false
	}
	var sCount, tCount [26]int
	for i := range s {
		sCount[s[i]-'a']++
	}
	for i := range t {
		tCount[t[i]-'a']++
	}
	return sCount == tCount
}
