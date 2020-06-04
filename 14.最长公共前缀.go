/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func check(strs []string, i int) bool {
	for j := range strs {
		if len(strs[j]) < i+1 {
			return true
		}
		if strs[j][i] != strs[0][i] {
			return true
		}
	}
	return false
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	i := 0
	for ; ; i++ {
		if check(strs, i) {
			break
		}
	}
	return strs[0][:i]
}

// @lc code=end

