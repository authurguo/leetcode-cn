/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start

func check(s []int) bool {
	l := len(s)
	i := 0
	j := l - 1
	for ; i < j; func() {
		i++
		j--
	}() {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func split(x int) []int {
	s := make([]int, 0, 10)
	y := 1
	z := (x / y) % 10
	for {
		if x/y == 0 {
			break
		}
		z = (x / y) % 10
		y = y * 10
		s = append(s, z)
	}
	return s
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	} else {
		return check(split(x))
	}
	return false
}

// @lc code=end

