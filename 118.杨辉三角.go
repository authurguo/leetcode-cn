/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */

// @lc code=start
func generate(numRows int) [][]int {
	t := make([][]int, 0, numRows)
	for i := 0; i < numRows; i++ {
		r := make([]int, 0, i+1)
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				r = append(r, 1)
			} else {
				k := t[i-1][j] + t[i-1][j-1]
				r = append(r, k)
			}
		}
		t = append(t, r)
	}
	return t
}

// @lc code=end

