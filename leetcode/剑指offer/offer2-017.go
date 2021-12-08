package offer

// https://leetcode-cn.com/problems/M1oyTv/
func minWindow(s string, t string) string {
	slen := len(s)
	tlen := len(t)
	if slen < tlen {
		return ""
	}
	//窗口内的元素种类及频率
	winCnt := make([]int, 128)
	//需要查找的元素种类及频率
	tCnt := make([]int, 128)
	for i := 0; i < tlen; i++ {
		tCnt[t[i]]++
	}
	//minLen需要找到一个不存在的值
	left, right, startIndex, minLen, instance := 0, 0, 0, slen+1, 0

	for right < slen {
		r := s[right]
		if tCnt[r] == 0 {
			right++
			continue
		}
		//当且仅当已有字符串目标字符出现的次数小于目标字符串字符的出现次数时，count才会+1
		//是为了后续能直接判断已有字符串是否已经包含了目标字符串的所有字符，不需要挨个比对字符出现的次数
		if winCnt[r] < tCnt[r] {
			instance++
		}
		winCnt[r]++
		right++
		for instance == tlen {
			if right-left < minLen {
				minLen = right - left
				startIndex = left
			}
			l := s[left]
			if tCnt[l] == 0 {
				left++
				continue
			}
			if winCnt[l] == tCnt[l] {
				instance--
			}
			winCnt[l]--
			left++
		}
	}
	if minLen == slen+1 {
		return ""
	}
	return s[startIndex : startIndex+minLen]
}
