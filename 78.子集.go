/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func deepCopy(s []int) []int {
	d := make([]int, len(s))
	for i, v := range s {
		d[i] = v
	}
	return d
}

func track(r [][]int, s []int) [][]int {
	if len(s) == 0 {
		return r
	}
	k := s[0]
	for _, v := range r {
		r = append(r, append(deepCopy(v), k))
	}
	return track(r, s[1:])
}

func subsets(nums []int) [][]int {
	r := make([][]int, 0, 128)
	r = append(r, []int{})
	return track(r, nums)
}

// @lc code=end

